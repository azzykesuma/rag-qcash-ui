// Package webui provides an embedded, offline dashboard for exported conversations.
package webui

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"llm-context-vault/pkg/sanitizer"
	"llm-context-vault/pkg/vault"
)

//go:embed index.html
var dashboard string

type session struct {
	File    string `json:"file"`
	Title   string `json:"title"`
	Tool    string `json:"tool"`
	Date    string `json:"date"`
	Project string `json:"project"`
}

// Run serves the dashboard on loopback until the process is stopped.
func Run(baseDir string, port int, openBrowser bool) error {
	if port < 0 || port > 65535 {
		return fmt.Errorf("port must be between 0 and 65535")
	}
	handler, err := NewHandler(baseDir)
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return err
	}
	defer listener.Close()
	url := "http://" + listener.Addr().String()
	fmt.Printf("Vault dashboard: %s\nVault: %s\nPress Ctrl+C to stop.\n", url, baseDir)
	if openBrowser {
		if err := launchBrowser(url); err != nil {
			fmt.Printf("Could not open browser: %v. Open the URL above manually.\n", err)
		}
	}
	server := &http.Server{Handler: handler, ReadHeaderTimeout: 5 * time.Second}
	return server.Serve(listener)
}

func launchBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	go func() { _ = cmd.Wait() }()
	return nil
}

// NewHandler exposes only exported Markdown and read-only vault operations.
func NewHandler(baseDir string) (http.Handler, error) {
	root, err := os.OpenRoot(filepath.Join(baseDir, "conversations", "markdown"))
	if err != nil {
		return nil, fmt.Errorf("open vault conversations (run vault scan first): %w", err)
	}
	// Validate the directory here; each request owns and closes its own root.
	root.Close()
	v := vault.New(baseDir, sanitizer.New(sanitizer.DefaultConfig()))
	mdDir := filepath.Join(baseDir, "conversations", "markdown")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, dashboard)
	})
	mux.HandleFunc("GET /api/sessions", func(w http.ResponseWriter, r *http.Request) {
		aliases, err := loadProjectAliases(baseDir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		entries, err := os.ReadDir(mdDir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		root, err := os.OpenRoot(mdDir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer root.Close()
		sessions := make([]session, 0)
		for _, entry := range entries {
			if !entry.Type().IsRegular() || filepath.Ext(entry.Name()) != ".md" {
				continue
			}
			f, err := root.Open(entry.Name())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			s := session{File: entry.Name(), Title: strings.TrimSuffix(entry.Name(), ".md"), Tool: "unknown"}
			scanner := bufio.NewScanner(f)
			for i := 0; i < 20 && scanner.Scan(); i++ {
				line := scanner.Text()
				switch {
				case strings.HasPrefix(line, "# "):
					s.Title = strings.TrimPrefix(line, "# ")
				case strings.HasPrefix(line, "- **Source Tool**:"):
					s.Tool = strings.Trim(strings.TrimPrefix(line, "- **Source Tool**:"), " `\r")
				case strings.HasPrefix(line, "- **Date**:"):
					s.Date = strings.Trim(strings.TrimPrefix(line, "- **Date**:"), " `\r")
				case strings.HasPrefix(line, "- **Project**:"):
					s.Project = strings.Trim(strings.TrimPrefix(line, "- **Project**:"), " `\r")
				case line == "---":
					i = 20
				}
			}
			f.Close()
			if alias := aliases[entry.Name()]; alias != "" {
				s.Project = alias
			}
			sessions = append(sessions, s)
		}
		writeJSON(w, sessions, nil)
	})
	mux.HandleFunc("GET /api/conversation", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("file")
		if name == "" || strings.ContainsAny(name, "/\\:") || filepath.Ext(name) != ".md" {
			http.Error(w, "Invalid conversation filename", http.StatusBadRequest)
			return
		}
		root, err := os.OpenRoot(mdDir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer root.Close()
		body, err := root.ReadFile(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write(body)
	})
	mux.HandleFunc("GET /api/search", func(w http.ResponseWriter, r *http.Request) {
		opts, _, err := queryOptions(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		results, err := v.SearchWithOptions(r.URL.Query().Get("q"), opts)
		writeJSON(w, results, err)
	})
	mux.HandleFunc("GET /api/context", func(w http.ResponseWriter, r *http.Request) {
		opts, maxChars, err := queryOptions(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		context, err := v.ContextWithOptions(r.URL.Query().Get("q"), opts, maxChars)
		writeJSON(w, context, err)
	})
	mux.HandleFunc("GET /api/audit", func(w http.ResponseWriter, r *http.Request) {
		violations, err := v.AuditAll()
		writeJSON(w, violations, err)
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'unsafe-inline'; style-src 'unsafe-inline'; frame-ancestors 'none'; base-uri 'none'")
		mux.ServeHTTP(w, r)
	}), nil
}

func loadProjectAliases(baseDir string) (map[string]string, error) {
	aliases := make(map[string]string)
	data, err := os.ReadFile(filepath.Join(baseDir, ".vault", "projects.json"))
	if os.IsNotExist(err) {
		return aliases, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &aliases); err != nil {
		return nil, fmt.Errorf("invalid .vault/projects.json: %w", err)
	}
	return aliases, nil
}

func queryOptions(r *http.Request) (vault.SearchOptions, int, error) {
	q := r.URL.Query()
	opts := vault.SearchOptions{Project: q.Get("project"), Tool: q.Get("tool"), After: q.Get("after"), Before: q.Get("before")}
	maxChars := 12000
	for name, target := range map[string]*int{"limit": &opts.Limit, "max-chars": &maxChars} {
		if value := q.Get(name); value != "" {
			parsed, err := strconv.Atoi(value)
			if err != nil || parsed < 1 {
				return opts, maxChars, fmt.Errorf("invalid %s", name)
			}
			*target = parsed
		}
	}
	if maxChars < 512 || maxChars > 1000000 {
		return opts, maxChars, fmt.Errorf("max-chars must be between 512 and 1000000")
	}
	return opts, maxChars, opts.Validate()
}

func writeJSON(w http.ResponseWriter, value any, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(value)
}
