package webui

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"llm-context-vault/pkg/vault"
)

func validSessionFile(name string) bool {
	return name != "" && name != ".md" && !strings.ContainsAny(name, "/\\:") && filepath.Base(name) == name && filepath.Ext(name) == ".md"
}

func validTopic(topic string) bool {
	if topic == "" || topic == "General development" {
		return true
	}
	for _, rule := range topicRules {
		if topic == rule.name {
			return true
		}
	}
	return false
}

func loadTopicOverrides(baseDir string) (map[string]string, error) {
	result := make(map[string]string)
	data, err := os.ReadFile(filepath.Join(baseDir, ".vault", "topics.json"))
	if os.IsNotExist(err) {
		return result, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid .vault/topics.json: %w", err)
	}
	return result, nil
}

func saveTopicOverrides(baseDir string, overrides map[string]string) error {
	dir := filepath.Join(baseDir, ".vault")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	f, err := os.CreateTemp(dir, "topics-*.tmp")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if err := f.Chmod(0600); err != nil {
		f.Close()
		return err
	}
	if err := json.NewEncoder(f).Encode(overrides); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), filepath.Join(dir, "topics.json"))
}

type relatedConversation struct {
	File    string `json:"file"`
	Title   string `json:"title"`
	Reason  string `json:"reason"`
	Snippet string `json:"snippet"`
}

type insights struct {
	Problem string                `json:"problem"`
	Finding string                `json:"finding"`
	Outcome string                `json:"outcome"`
	Related []relatedConversation `json:"related"`
}

var summaryMarkup = regexp.MustCompile("(?m)^[#>*`-]+\\s*|<[^>]+>")
var summaryWords = regexp.MustCompile(`[\pL\pN][\pL\pN_-]*`)
var confirmedOutcome = regexp.MustCompile(`(?i)\b(it works|that works|works now|fixed now|issue resolved|problem solved|that fixed it|it's fixed|thank you,? (?:it works|that worked))\b`)
var summaryStop = map[string]bool{"please": true, "could": true, "would": true, "this": true, "that": true, "with": true, "from": true, "what": true, "when": true, "where": true, "have": true, "about": true, "there": true, "your": true, "just": true, "check": true, "help": true, "issue": true, "error": true, "fix": true, "whole": true, "generate": true, "make": true, "tell": true, "want": true, "need": true, "repo": true, "code": true, "file": true, "work": true}

func shortText(text string, limit int) string {
	text = strings.Join(strings.Fields(summaryMarkup.ReplaceAllString(text, " ")), " ")
	r := []rune(text)
	if len(r) > limit {
		return string(r[:limit]) + "…"
	}
	return text
}

func conversationInsights(mdDir, name string, v *vault.Vault) (insights, error) {
	root, err := os.OpenRoot(mdDir)
	if err != nil {
		return insights{}, err
	}
	defer root.Close()
	f, err := root.Open(name)
	if err != nil {
		return insights{}, err
	}
	defer f.Close()
	result := insights{Outcome: "Unknown — no confirmed resolution in this session."}
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	role, inTools, firstUserDone, assistantSeen := "", false, false, false
	title := strings.TrimSuffix(name, ".md")
	project := ""
	var userText, assistantText strings.Builder
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "# ") && role == "" {
			title = strings.TrimPrefix(line, "# ")
		}
		if strings.HasPrefix(line, "- **Project**:") && role == "" {
			project = strings.Trim(strings.TrimPrefix(line, "- **Project**:"), " `\r")
		}
		if strings.HasPrefix(line, "## Turn ") {
			if strings.HasSuffix(line, ": User") {
				role = "user"
				if userText.Len() > 0 {
					firstUserDone = true
				}
			} else if strings.HasSuffix(line, ": Assistant") {
				role = "assistant"
				assistantSeen = true
				assistantText.Reset()
			} else {
				role = ""
			}
			inTools = false
			continue
		}
		if strings.HasPrefix(line, "<details>") {
			inTools = true
			continue
		}
		if strings.HasPrefix(line, "</details>") {
			inTools = false
			continue
		}
		if inTools || line == "---" || line == "" || strings.HasPrefix(line, "```") || strings.HasPrefix(line, "-") {
			continue
		}
		if role == "user" && !firstUserDone && userText.Len() < 450 {
			userText.WriteString(line + " ")
		}
		if role == "user" && assistantSeen && confirmedOutcome.MatchString(line) {
			result.Outcome = "Confirmed by user: " + shortText(line, 180)
		}
		if role == "assistant" && assistantText.Len() < 600 {
			assistantText.WriteString(line + " ")
		}
	}
	if err := scanner.Err(); err != nil {
		return insights{}, err
	}
	result.Problem = shortText(userText.String(), 300)
	result.Finding = shortText(assistantText.String(), 300)
	if result.Problem == "" {
		result.Problem = "No user problem statement found."
	}
	if result.Finding == "" {
		result.Finding = "No assistant finding recorded."
	}
	// Suggestions are based on the problem/title, not an unverified assistant answer.
	terms, seen := make([]string, 0, 8), make(map[string]bool)
	for _, word := range summaryWords.FindAllString(strings.ToLower(title+" "+result.Problem), -1) {
		if len(word) < 4 || summaryStop[word] || seen[word] {
			continue
		}
		seen[word] = true
		terms = append(terms, word)
		if len(terms) == 8 {
			break
		}
	}
	if len(terms) == 0 {
		return result, nil
	}
	matches, err := v.SearchWithOptions(strings.Join(terms, " "), vault.SearchOptions{Limit: 50})
	if err != nil {
		return insights{}, err
	}
	type candidate struct {
		relatedConversation
		score int
	}
	var ranked []candidate
	for _, match := range matches {
		file := filepath.Base(match.Path)
		if file == name || !validSessionFile(file) {
			continue
		}
		titleWords, bodyWords := make(map[string]bool), make(map[string]bool)
		for _, word := range summaryWords.FindAllString(strings.ToLower(match.Title), -1) {
			titleWords[word] = true
		}
		for _, word := range summaryWords.FindAllString(strings.ToLower(match.Snippet), -1) {
			bodyWords[word] = true
		}
		var shared []string
		score, titleHits := 0, 0
		for _, term := range terms {
			if titleWords[term] {
				score += 3
				titleHits++
				shared = append(shared, term)
			} else if bodyWords[term] {
				score++
				shared = append(shared, term)
			}
		}
		if len(shared) < 2 || (titleHits == 0 && len(shared) < 3) {
			continue
		}
		if project != "" && strings.EqualFold(project, match.Project) {
			score += 2
		}
		ranked = append(ranked, candidate{relatedConversation{File: file, Title: match.Title, Reason: "Shared terms: " + strings.Join(shared, ", "), Snippet: shortText(match.Snippet, 160)}, score})
	}
	sort.SliceStable(ranked, func(i, j int) bool { return ranked[i].score > ranked[j].score })
	for _, match := range ranked[:min(5, len(ranked))] {
		result.Related = append(result.Related, match.relatedConversation)
	}
	return result, nil
}
