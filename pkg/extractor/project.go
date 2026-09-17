package extractor

import (
	"os"
	"path/filepath"
	"strings"
)

func projectName(path string) string {
	path = strings.TrimRight(strings.ReplaceAll(path, "\\", "/"), "/")
	if path == "" {
		return ""
	}
	if home, err := os.UserHomeDir(); err == nil && strings.EqualFold(path, strings.TrimRight(filepath.ToSlash(home), "/")) {
		return ""
	}
	parts := strings.Split(path, "/")
	name := parts[len(parts)-1]
	if len(parts) >= 2 {
		parent := strings.ToLower(parts[len(parts)-2])
		if parent == "users" || parent == "home" {
			return ""
		}
	}
	if name == "." || name == ".." || strings.Contains(name, ":") {
		return ""
	}
	return name
}
