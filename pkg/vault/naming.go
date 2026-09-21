package vault

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"llm-context-vault/pkg/models"
)

type existingExport struct {
	Slug      string
	CreatedAt time.Time
}

// Recover naming decisions from exported IDs, including legacy filenames. This
// registry is rebuildable; deleting a local cache must never rename old exports.
func (v *Vault) loadExportNames() error {
	if v.exportNames != nil {
		return nil
	}
	v.exportNames = make(map[string]existingExport)
	v.identityAliases = make(map[string]string)
	entries, err := os.ReadDir(filepath.Join(v.BaseDir, "conversations", "markdown"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		v.exportNames = nil
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		f, err := os.Open(filepath.Join(v.BaseDir, "conversations", "markdown", entry.Name()))
		if err != nil {
			v.exportNames = nil
			return err
		}
		scanner := bufio.NewScanner(f)
		scanner.Buffer(make([]byte, 4096), 1024*1024)
		var id string
		existing := existingExport{Slug: strings.TrimSuffix(entry.Name(), ".md")}
		for i := 0; i < 12 && scanner.Scan(); i++ {
			line := scanner.Text()
			if strings.HasPrefix(line, "- **ID**:") {
				id = strings.Trim(strings.TrimPrefix(line, "- **ID**:"), " `\r")
			}
			if strings.HasPrefix(line, "- **Date**:") {
				existing.CreatedAt, _ = time.ParseInLocation("2006-01-02 15:04:05", strings.Trim(strings.TrimPrefix(line, "- **Date**:"), " `\r"), time.Local)
			}
		}
		err = scanner.Err()
		f.Close()
		if err != nil {
			v.exportNames = nil
			return fmt.Errorf("read export header: %w", err)
		}
		if id != "" {
			if _, found := v.exportNames[id]; !found {
				v.exportNames[id] = existing
			}
		}
	}
	if v.rebuildingDataset {
		return nil
	}
	// Older AGY exports hashed language names in random map order. Match their
	// canonical content back to the published ID rather than renaming/reimporting
	// them. Also recover exact timestamps (Markdown only retains whole seconds).
	dataset, err := os.Open(filepath.Join(v.BaseDir, "conversations", "dataset.jsonl"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		v.exportNames = nil
		return err
	}
	defer dataset.Close()
	scanner := bufio.NewScanner(dataset)
	scanner.Buffer(make([]byte, 64*1024), 64*1024*1024)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		var conv models.Conversation
		if err := json.Unmarshal(scanner.Bytes(), &conv); err != nil {
			v.exportNames = nil
			return fmt.Errorf("read existing dataset: %w", err)
		}
		if existing, ok := v.exportNames[conv.ID]; ok {
			if !conv.CreatedAt.IsZero() {
				existing.CreatedAt = conv.CreatedAt
				v.exportNames[conv.ID] = existing
			}
			for _, candidate := range []*models.Conversation{&conv, v.Sanitizer.SanitizeConversation(&conv)} {
				canonicalID := stableConversationID(candidate)
				if _, found := v.identityAliases[canonicalID]; !found {
					v.identityAliases[canonicalID] = conv.ID
				}
				project := candidate.Project
				candidate.Project = ""
				legacyID := stableConversationID(candidate)
				candidate.Project = project
				if _, found := v.identityAliases[legacyID]; !found {
					v.identityAliases[legacyID] = conv.ID
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		v.exportNames = nil
		return err
	}
	return nil
}

var topicNoise = regexp.MustCompile("(?i)https?://\\S+|`[^`]*`|<[^>]*>|\\[[A-Z_]*REDACTED[A-Z_]*\\]")
var topicWords = regexp.MustCompile(`[\pL][\pL\pN_-]*`)
var namingStop = wordSet("a an the and or is are i my to in of it we you your this that these those can could would should please help check look want need have has had do does did how why what when where with for from on at by as be been being me us our they their them then now here there also just some all any not no yes hello hi hey thanks thank okay ok lets let use using used run running command tool output input user assistant system response message new untitled code file files implementation implement result results issue problem error fix fixed resolve resolved debug investigate add update remove create build test tests failed failure task work working will into about after before more make get got see find follow instructions environment context skills permissions summary text content details turn")

func wordSet(s string) map[string]bool {
	m := make(map[string]bool)
	for _, w := range strings.Fields(s) {
		m[w] = true
	}
	return m
}

func topicTerms(s string) []string {
	s = topicNoise.ReplaceAllString(s, " ")
	var terms []string
	for _, w := range topicWords.FindAllString(strings.ToLower(s), -1) {
		w = strings.Trim(w, "-_")
		if len(w) < 3 || len(w) > 28 || namingStop[w] {
			continue
		}
		digits := 0
		for _, r := range w {
			if unicode.IsDigit(r) {
				digits++
			}
		}
		if digits > 3 || strings.Contains(w, "redacted") {
			continue
		}
		terms = append(terms, w)
	}
	return terms
}

// conversationTopic is intentionally extractive: it names recurring subjects,
// without claiming that a requested fix actually succeeded. Work is bounded per
// message and sampled evenly across long sessions, including their final turns.
func conversationTopic(conv *models.Conversation) string {
	freq := make(map[string]int)
	var samples []string
	step := max(1, (len(conv.Messages)+127)/128)
	for i, msg := range conv.Messages {
		if i%step != 0 && i != len(conv.Messages)-1 {
			continue
		}
		if msg.Role != "user" && msg.Role != "assistant" {
			continue
		}
		text := msg.Content
		if strings.Contains(text, "<environment_context>") || strings.HasPrefix(text, "# AGENTS.md") || strings.Contains(text, "<system-reminder>") {
			continue
		}
		text = truncateRunes(text, 3000)
		samples = append(samples, text)
		seen := make(map[string]bool)
		for _, term := range topicTerms(text) {
			if !seen[term] {
				freq[term]++
				seen[term] = true
			}
		}
	}
	// A genuine editor-assigned title is usually better than a bag of keywords.
	titleTerms := topicTerms(conv.Title)
	titleLower := strings.ToLower(conv.Title)
	if len(titleTerms) >= 2 && len(titleTerms) <= 8 && len([]rune(conv.Title)) <= 90 && !strings.Contains(titleLower, "http") && !strings.Contains(titleLower, "redacted") && !strings.ContainsAny(conv.Title, "<>={}\\/") {
		supported := 0
		for _, term := range titleTerms {
			if freq[term] >= 2 {
				supported++
			}
		}
		if supported >= 2 {
			return safeTopicSlug(strings.Join(titleTerms, "-"))
		}
	}
	// Pick a short contiguous topic phrase supported across multiple messages.
	best, bestScore := "", 0
	for _, sample := range samples {
		for _, line := range strings.Split(sample, "\n") {
			if strings.HasPrefix(strings.TrimSpace(line), "```") || strings.Contains(line, "https://") {
				continue
			}
			terms := topicTerms(line)
			for start := range terms {
				for size := 2; size <= 4 && start+size <= len(terms); size++ {
					phrase := terms[start : start+size]
					score := 0
					seen := make(map[string]bool)
					for _, term := range phrase {
						if !seen[term] {
							score += min(freq[term], 8)
							seen[term] = true
						}
					}
					score = score*10/size + min(size, 3)
					candidate := strings.Join(phrase, "-")
					if score > bestScore || score == bestScore && candidate < best {
						best, bestScore = candidate, score
					}
				}
			}
		}
	}
	if best == "" {
		terms := make([]string, 0, len(freq))
		for term := range freq {
			terms = append(terms, term)
		}
		sort.Slice(terms, func(i, j int) bool {
			if freq[terms[i]] != freq[terms[j]] {
				return freq[terms[i]] > freq[terms[j]]
			}
			return terms[i] < terms[j]
		})
		best = strings.Join(terms[:min(3, len(terms))], "-")
	}
	return safeTopicSlug(best)
}

func safeTopicSlug(s string) string {
	s = strings.Trim(toKebabCase(s, 1000), "-")
	s = strings.Trim(truncateRunes(s, 60), "-")
	if s == "" {
		return "conversation"
	}
	return s
}

func truncateRunes(s string, limit int) string {
	runes := []rune(s)
	if len(runes) > limit {
		return string(runes[:limit])
	}
	return s
}
