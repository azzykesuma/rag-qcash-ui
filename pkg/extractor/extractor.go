package extractor

import (
	"llm-context-vault/pkg/models"
)

// Extractor defines the interface for tool-specific conversation log parsers
type Extractor interface {
	Name() string
	CanHandle(path string) bool
	Extract(path string) (*models.Conversation, error)
}
