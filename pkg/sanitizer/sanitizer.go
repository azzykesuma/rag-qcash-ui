package sanitizer

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"llm-context-vault/pkg/models"
)

// Rule represents a redaction regex rule with replacement text
type Rule struct {
	Name        string
	Pattern     *regexp.Regexp
	Replacement string
}

// Config controls which sanitization features are active
type Config struct {
	RedactSecrets   bool     `json:"redact_secrets"`
	RedactPaths     bool     `json:"redact_paths"`
	RedactEmails    bool     `json:"redact_emails"`
	RedactIPs       bool     `json:"redact_ips"`
	CustomKeywords  []string `json:"custom_keywords"`  // Custom terms to redact (e.g. company names)
	RedactionNotice string   `json:"redaction_notice"` // Text to replace custom keywords with
}

// DefaultConfig returns recommended sanitization settings
func DefaultConfig() Config {
	return Config{
		RedactSecrets:   true,
		RedactPaths:     true,
		RedactEmails:    true,
		RedactIPs:       true,
		CustomKeywords:  []string{},
		RedactionNotice: "[REDACTED]",
	}
}

// Sanitizer provides methods to strip sensitive information from messages and conversations
type Sanitizer struct {
	cfg          Config
	secretRules  []Rule
	pathRules    []Rule
	emailRule    *regexp.Regexp
	ipRule       *regexp.Regexp
	keywordRules []Rule
}

// New creates a new Sanitizer instance with precompiled regex patterns
func New(cfg Config) *Sanitizer {
	s := &Sanitizer{
		cfg: cfg,
	}
	s.initRules()
	return s
}

func (s *Sanitizer) initRules() {
	// 1. Secret & Key patterns
	s.secretRules = []Rule{
		{
			Name:        "OpenAI API Key",
			Pattern:     regexp.MustCompile(`\b(sk-(?:proj-)?[a-zA-Z0-9_\-]{20,})\b`),
			Replacement: "[OPENAI_API_KEY_REDACTED]",
		},
		{
			Name:        "Anthropic API Key",
			Pattern:     regexp.MustCompile(`\b(sk-ant-[a-zA-Z0-9_\-]{20,})\b`),
			Replacement: "[ANTHROPIC_API_KEY_REDACTED]",
		},
		{
			Name:        "Google API Key",
			Pattern:     regexp.MustCompile(`\b(AIza[0-9A-Za-z\-_]{35})\b`),
			Replacement: "[GOOGLE_API_KEY_REDACTED]",
		},
		{
			Name:        "GitHub Token",
			Pattern:     regexp.MustCompile(`\b(gh[pousr]_[A-Za-z0-9_]{36,}|github_pat_[0-9a-zA-Z_]{82})\b`),
			Replacement: "[GITHUB_TOKEN_REDACTED]",
		},
		{
			Name:        "AWS Access Key ID",
			Pattern:     regexp.MustCompile(`\b(AKIA[0-9A-Z]{16})\b`),
			Replacement: "[AWS_ACCESS_KEY_REDACTED]",
		},
		{
			Name:        "AWS Secret Access Key",
			Pattern:     regexp.MustCompile(`(?i)(aws_secret_access_key|aws_sec_key|aws_secret)\s*[:=]\s*['"]?([A-Za-z0-9\/+=]{40})['"]?`),
			Replacement: "$1: [AWS_SECRET_REDACTED]",
		},
		{
			Name:        "Private Keys",
			Pattern:     regexp.MustCompile(`-----BEGIN [A-Z ]*PRIVATE KEY-----[\s\S]*?-----END [A-Z ]*PRIVATE KEY-----`),
			Replacement: "[PRIVATE_KEY_BLOCK_REDACTED]",
		},
		{
			Name:        "JWT / Bearer Token",
			Pattern:     regexp.MustCompile(`(?i)\bBearer\s+(ey[A-Za-z0-9\-_=]+\.ey[A-Za-z0-9\-_=]+\.[A-Za-z0-9\-_=+/]+)`),
			Replacement: "Bearer [JWT_TOKEN_REDACTED]",
		},
		{
			Name:        "Database Connection Strings",
			Pattern:     regexp.MustCompile(`(?i)\b((?:postgres|postgresql|mysql|mongodb|redis|amqp|mssql):\/\/)([^:\s]+):([^@\s\[\]]+)@`),
			Replacement: "$1$2:[PASSWORD_REDACTED]@",
		},
		{
			Name:        "Bitbucket / Atlassian Server PAT",
			Pattern:     regexp.MustCompile(`(?i)\b(?:[A-Za-z0-9+/=]{10,}:[A-Za-z0-9+/=]{20,}|(?:BBDC-|squ_|ctx7sk-|jam_pat_)[A-Za-z0-9_\-\+\/=]{20,}|(?:MT[0-9A-Za-z]{2}|MD[0-9A-Za-z]{2}|OT[0-9A-Za-z]{2}|Mz[0-9A-Za-z]{2}|Nz[0-9A-Za-z]{2}|Mj[0-9A-Za-z]{2})[A-Za-z0-9\+\/=]{25,})`),
			Replacement: "[ATLASSIAN_TOKEN_REDACTED]",
		},
		{
			Name:        "Token Prompt Assignment",
			Pattern:     regexp.MustCompile(`(?i)\b(use this token|token is|bearer token|my token is|api token|access token|using token|using this(?: [a-zA-Z0-9_\-]+)* token)\s*[:=]?\s*([A-Za-z0-9\+\/=_]{20,})`),
			Replacement: "$1: [TOKEN_REDACTED]",
		},
		{
			Name:        "SonarQube Token",
			Pattern:     regexp.MustCompile(`\b(squ_[0-9a-fA-F]{30,})\b`),
			Replacement: "[SONARQUBE_TOKEN_REDACTED]",
		},
		{
			Name:        "Jam PAT",
			Pattern:     regexp.MustCompile(`\b(jam_pat_[A-Za-z0-9_\-]{20,})\b`),
			Replacement: "[JAM_PAT_REDACTED]",
		},
		{
			Name:        "Context7 Key",
			Pattern:     regexp.MustCompile(`\b(ctx7sk-[A-Za-z0-9_\-]{20,})\b`),
			Replacement: "[CONTEXT7_KEY_REDACTED]",
		},
		{
			Name:        "Generic Secret Assignment",
			Pattern:     regexp.MustCompile(`(?i)\b([a-zA-Z0-9_\-]*(?:token|secret|api[_-]?key|access[_-]?key|pat|auth|password|passwd|pwd|credentials))\s*([:=])\s*['"]([A-Za-z0-9_\-\.\+\/=]{8,})['"]`),
			Replacement: "$1$2 \"[SECRET_REDACTED]\"",
		},
	}

	// 2. Machine & User Paths & Internal Domains
	s.pathRules = []Rule{
		{
			Name:        "Windows User Directory",
			Pattern:     regexp.MustCompile(`(?i)[a-zA-Z]:\\Users\\[a-zA-Z0-9_\-\.]+`),
			Replacement: "~",
		},
		{
			Name:        "Unix / Mac User Directory",
			Pattern:     regexp.MustCompile(`\/(?:home|Users)\/[a-zA-Z0-9_\-\.]+`),
			Replacement: "~",
		},
		{
			Name:        "Internal Corporate Domain",
			Pattern:     regexp.MustCompile(`(?i)\b[a-zA-Z0-9_\-\.]*\.bri\.co\.id\b`),
			Replacement: "internal-service.example.com",
		},
	}

	// 3. Email Rule
	s.emailRule = regexp.MustCompile(`\b[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}\b`)

	// 4. IP Address Rule
	s.ipRule = regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)

	// 5. Custom keywords
	s.keywordRules = make([]Rule, 0, len(s.cfg.CustomKeywords))
	for _, kw := range s.cfg.CustomKeywords {
		if strings.TrimSpace(kw) == "" {
			continue
		}
		escaped := regexp.QuoteMeta(strings.TrimSpace(kw))
		s.keywordRules = append(s.keywordRules, Rule{
			Name:        "Custom Keyword: " + kw,
			Pattern:     regexp.MustCompile(`(?i)\b` + escaped + `\b`),
			Replacement: s.cfg.RedactionNotice,
		})
	}
}

// SanitizeText takes any text and applies all enabled redaction rules
func (s *Sanitizer) SanitizeText(text string) string {
	if text == "" {
		return ""
	}

	result := text

	// 1. Secrets
	if s.cfg.RedactSecrets {
		for _, rule := range s.secretRules {
			result = rule.Pattern.ReplaceAllString(result, rule.Replacement)
		}
	}

	// 2. System Paths
	if s.cfg.RedactPaths {
		for _, rule := range s.pathRules {
			result = rule.Pattern.ReplaceAllString(result, rule.Replacement)
		}
	}

	// 3. Emails
	if s.cfg.RedactEmails {
		result = s.emailRule.ReplaceAllString(result, "[REDACTED_EMAIL]")
	}

	// 4. IP Addresses (preserve localhost / loopback / zero)
	if s.cfg.RedactIPs {
		result = s.ipRule.ReplaceAllStringFunc(result, func(match string) string {
			parsed := net.ParseIP(match)
			if parsed == nil {
				return match
			}
			// Keep local loopback and zero addresses as they are standard in code examples
			if parsed.IsLoopback() || parsed.IsUnspecified() || match == "127.0.0.1" || match == "0.0.0.0" {
				return match
			}
			return "[REDACTED_IP]"
		})
	}

	// 5. Custom Keywords
	for _, rule := range s.keywordRules {
		result = rule.Pattern.ReplaceAllString(result, rule.Replacement)
	}

	return result
}

// SanitizeConversation processes an entire conversation in-place or returns a sanitized clone
func (s *Sanitizer) SanitizeConversation(conv *models.Conversation) *models.Conversation {
	if conv == nil {
		return nil
	}

	cloned := *conv
	cloned.Title = s.SanitizeText(conv.Title)
	cloned.Description = s.SanitizeText(conv.Description)

	sanitizedMessages := make([]models.Message, len(conv.Messages))
	for i, msg := range conv.Messages {
		sanitizedMsg := msg
		sanitizedMsg.Content = s.SanitizeText(msg.Content)

		if len(msg.ToolCalls) > 0 {
			sanitizedToolCalls := make([]models.ToolCall, len(msg.ToolCalls))
			for j, tc := range msg.ToolCalls {
				sanitizedTc := tc
				sanitizedTc.Arguments = s.SanitizeText(tc.Arguments)
				sanitizedTc.Summary = s.SanitizeText(tc.Summary)
				sanitizedTc.Output = s.SanitizeText(tc.Output)
				sanitizedToolCalls[j] = sanitizedTc
			}
			sanitizedMsg.ToolCalls = sanitizedToolCalls
		}

		sanitizedMessages[i] = sanitizedMsg
	}

	cloned.Messages = sanitizedMessages
	return &cloned
}

// AuditText checks if any sensitive pattern is detected in the text and returns matching rule names
func (s *Sanitizer) AuditText(text string) []string {
	var matches []string

	for _, rule := range s.secretRules {
		if rule.Pattern.MatchString(text) {
			matches = append(matches, fmt.Sprintf("Secret [%s]", rule.Name))
		}
	}

	for _, rule := range s.pathRules {
		if rule.Pattern.MatchString(text) {
			matches = append(matches, fmt.Sprintf("Path [%s]", rule.Name))
		}
	}

	if s.emailRule.MatchString(text) {
		matches = append(matches, "Email Address")
	}

	return matches
}
