package exporter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"llm-context-vault/pkg/models"
)

// MarkdownExporter writes conversations as GitHub-flavored Markdown
type MarkdownExporter struct{}

func NewMarkdownExporter() *MarkdownExporter {
	return &MarkdownExporter{}
}

func (e *MarkdownExporter) Export(conv *models.Conversation, outPath string) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	var b strings.Builder

	// Header & Metadata Frontmatter
	b.WriteString(fmt.Sprintf("# %s\n\n", conv.Title))
	b.WriteString(fmt.Sprintf("- **ID**: `%s`\n", conv.ID))
	b.WriteString(fmt.Sprintf("- **Source Tool**: `%s`\n", conv.SourceTool))
	b.WriteString(fmt.Sprintf("- **Date**: `%s`\n", conv.CreatedAt.Format("2006-01-02 15:04:05")))

	if len(conv.Languages) > 0 {
		b.WriteString(fmt.Sprintf("- **Languages**: %s\n", strings.Join(conv.Languages, ", ")))
	}
	if len(conv.Tags) > 0 {
		b.WriteString(fmt.Sprintf("- **Tags**: `%s`\n", strings.Join(conv.Tags, "`, `")))
	}

	b.WriteString("\n---\n\n")

	// Messages
	for i, msg := range conv.Messages {
		switch strings.ToLower(msg.Role) {
		case "user":
			b.WriteString(fmt.Sprintf("## Turn %d: User\n\n", i+1))
			b.WriteString(msg.Content)
			b.WriteString("\n\n")

		case "assistant":
			b.WriteString(fmt.Sprintf("## Turn %d: Assistant\n\n", i+1))

			if len(msg.ToolCalls) > 0 {
				b.WriteString("<details><summary>🔧 Tool Invocations (" + fmt.Sprintf("%d", len(msg.ToolCalls)) + ")</summary>\n\n")
				for _, tc := range msg.ToolCalls {
					b.WriteString(fmt.Sprintf("- **%s**: %s\n", tc.Name, tc.Summary))
				}
				b.WriteString("\n</details>\n\n")
			}

			b.WriteString(msg.Content)
			b.WriteString("\n\n")

		case "system":
			b.WriteString(fmt.Sprintf("> **System Instruction**:\n> %s\n\n", strings.ReplaceAll(msg.Content, "\n", "\n> ")))
		}

		b.WriteString("---\n\n")
	}

	return os.WriteFile(outPath, []byte(b.String()), 0644)
}
