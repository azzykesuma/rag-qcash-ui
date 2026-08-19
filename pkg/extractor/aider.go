package extractor

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"llm-context-vault/pkg/models"
)

// AiderExtractor extracts conversations from Aider chat history files
type AiderExtractor struct{}

func NewAiderExtractor() *AiderExtractor {
	return &AiderExtractor{}
}

func (e *AiderExtractor) Name() string {
	return "aider"
}

func (e *AiderExtractor) CanHandle(path string) bool {
	base := strings.ToLower(filepath.Base(path))
	return strings.Contains(base, ".aider.chat.history.md") || strings.Contains(base, "aider")
}

func (e *AiderExtractor) Extract(targetPath string) (*models.Conversation, error) {
	file, err := os.Open(targetPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open aider history: %w", err)
	}
	defer file.Close()

	var messages []models.Message
	var currentRole string
	var currentLines []string
	var firstUserInput string

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	flushTurn := func() {
		if currentRole != "" && len(currentLines) > 0 {
			content := strings.TrimSpace(strings.Join(currentLines, "\n"))
			if content != "" {
				messages = append(messages, models.Message{
					Role:    currentRole,
					Content: content,
				})
			}
			currentLines = nil
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#### USER") || strings.HasPrefix(line, "> ") {
			flushTurn()
			currentRole = "user"
			trimmed := strings.TrimPrefix(line, "#### USER")
			trimmed = strings.TrimPrefix(trimmed, "> ")
			if trimmed != "" {
				currentLines = append(currentLines, trimmed)
			}
			continue
		} else if strings.HasPrefix(line, "#### ASSISTANT") {
			flushTurn()
			currentRole = "assistant"
			continue
		}

		if currentRole == "user" && firstUserInput == "" && strings.TrimSpace(line) != "" {
			firstUserInput = strings.TrimSpace(line)
		}

		currentLines = append(currentLines, line)
	}

	flushTurn()

	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages extracted from aider history %s", targetPath)
	}

	title := generateTitle(firstUserInput)
	if title == "Untitled Conversation" {
		title = "Aider Session - " + filepath.Base(filepath.Dir(targetPath))
	}

	return &models.Conversation{
		ID:         fmt.Sprintf("aider_%d", time.Now().Unix()),
		SourceTool: "aider",
		Title:      title,
		CreatedAt:  time.Now(),
		Tags:       []string{"aider", "terminal", "coding"},
		Messages:   messages,
	}, nil
}
