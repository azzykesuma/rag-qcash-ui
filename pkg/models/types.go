package models

import (
	"time"
)

// Message represents a single turn in a conversation
type Message struct {
	Role      string         `json:"role"`                 // "user", "assistant", "system", "tool"
	Content   string         `json:"content"`              // Text content of the message
	Timestamp *time.Time     `json:"timestamp,omitempty"`  // Timestamp of the turn
	ToolCalls []ToolCall     `json:"tool_calls,omitempty"` // Associated tool calls if any
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// ToolCall represents a tool invocation and its sanitized summary
type ToolCall struct {
	Name      string `json:"name"`
	Summary   string `json:"summary,omitempty"`
	Arguments string `json:"arguments,omitempty"`
	Output    string `json:"output,omitempty"`
}

// Conversation represents a normalized session from any coding assistant
type Conversation struct {
	ID          string            `json:"id"`
	SourceTool  string            `json:"source_tool"` // "antigravity", "opencode", "aider", "cursor", "codex", etc.
	Title       string            `json:"title"`
	Description string            `json:"description,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	Languages   []string          `json:"languages,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Messages    []Message         `json:"messages"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// ShareGPTFormat represents the standard ShareGPT / HuggingFace dataset format
type ShareGPTFormat struct {
	ID            string            `json:"id"`
	Conversations []ShareGPTMessage `json:"conversations"`
	Source        string            `json:"source,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
}

type ShareGPTMessage struct {
	From  string `json:"from"` // "human", "gpt", "system"
	Value string `json:"value"`
}
