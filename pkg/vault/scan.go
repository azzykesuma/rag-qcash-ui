package vault

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"llm-context-vault/pkg/exporter"
	"llm-context-vault/pkg/models"
)

// Bump when extraction, filtering, or export behavior changes.
const scanVersion = "3"

type ScanOptions struct {
	Full     bool
	Workers  int
	Progress func(string)
}

type ScanSource struct{ Tool, Path string }

type fileStamp struct {
	Size     int64 `json:"size"`
	Modified int64 `json:"modified"`
}

type scanEntry struct {
	Revision       string               `json:"revision"`
	ConversationID string               `json:"conversation_id,omitempty"`
	Exports        map[string]fileStamp `json:"exports,omitempty"`
}

type scanState struct {
	Version string               `json:"version"`
	Dataset fileStamp            `json:"dataset"`
	Entries map[string]scanEntry `json:"entries"`
}

func stamp(path string) (fileStamp, error) {
	info, err := os.Stat(path)
	if err != nil {
		return fileStamp{}, err
	}
	return fileStamp{Size: info.Size(), Modified: info.ModTime().UnixNano()}, nil
}

func sourceKey(tool, path, session string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(tool+"\x00"+path+"\x00"+session)))
}

func (v *Vault) lockWriter() (func(), error) {
	dir := filepath.Join(v.BaseDir, ".vault")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, err
	}
	path := filepath.Join(dir, "write.lock")
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return nil, fmt.Errorf("vault writer lock %s: %w; if a previous process crashed, remove the lock only after confirming it is no longer running", path, err)
	}
	fmt.Fprintf(f, "pid=%d\n", os.Getpid())
	f.Close()
	return func() { _ = os.Remove(path) }, nil
}

// AcquireWriterLock serializes commands that read and then modify vault exports.
func (v *Vault) AcquireWriterLock() (func(), error) { return v.lockWriter() }

func (v *Vault) progress(message string) {
	if v.ScanOptions.Progress != nil {
		v.ScanOptions.Progress(message)
	}
}

func (v *Vault) scanOne(tool, path string) (int, int, []string, error) {
	reports, err := v.ScanSources([]ScanSource{{Tool: tool, Path: path}})
	if len(reports) == 0 {
		return 0, 0, nil, err
	}
	r := reports[0]
	if err == nil && r.FailedCount > 0 {
		err = fmt.Errorf("%s scan failed for %d source/session(s)", tool, r.FailedCount)
	}
	return r.ImportedCount, r.SkippedCount, r.Warnings, err
}

// ScanSources is shared by unified and tool-specific scans. It holds one writer
// lock, batches dataset writes across sources, and checkpoints only after flush.
func (v *Vault) ScanSources(sources []ScanSource) ([]ToolScanReport, error) {
	release, err := v.lockWriter()
	if err != nil {
		return nil, err
	}
	defer release()
	dataset := filepath.Join(v.BaseDir, "conversations", "dataset.jsonl")
	cachePath := filepath.Join(v.BaseDir, ".vault", "scan-state.json")
	version := scanVersion + ":" + v.Sanitizer.Fingerprint()
	state := scanState{Version: version, Entries: make(map[string]scanEntry)}
	data, readErr := os.ReadFile(cachePath)
	if readErr == nil {
		if err := json.Unmarshal(data, &state); err != nil {
			v.progress("Scan cache is invalid; rebuilding.")
			state = scanState{}
		}
	} else if !os.IsNotExist(readErr) {
		return nil, readErr
	}
	currentDataset, _ := stamp(dataset)
	invalidated := state.Version != version || state.Dataset != currentDataset || state.Entries == nil
	if state.Entries == nil {
		state.Entries = make(map[string]scanEntry)
	}
	if invalidated {
		// Preserve source ownership so successful retries can retire old exports,
		// but force every entry to be processed before it becomes current again.
		for key, entry := range state.Entries {
			entry.Revision = ""
			state.Entries[key] = entry
		}
	}
	// A malformed cache can contain partially decoded entries; never trust it.
	if readErr == nil && !json.Valid(data) {
		state.Entries = make(map[string]scanEntry)
	}
	state.Version = version
	v.exportNames = nil
	v.retireEntries = nil
	v.jsonExport.BeginBatch()
	var reports []ToolScanReport
	seenSources := make(map[string]bool)
	for _, source := range sources {
		path, err := filepath.Abs(source.Path)
		if err != nil {
			return reports, err
		}
		source.Path = filepath.Clean(path)
		key := sourceKey(source.Tool, source.Path, "")
		if seenSources[key] {
			continue
		}
		seenSources[key] = true
		report := v.scanSource(source, &state)
		reports = append(reports, report)
	}
	referenced := make(map[string]bool)
	for _, entry := range state.Entries {
		if entry.ConversationID != "" {
			referenced[entry.ConversationID] = true
		}
	}
	for _, entry := range v.retireEntries {
		if !referenced[entry.ConversationID] {
			v.jsonExport.DeleteJSONL(entry.ConversationID)
		}
	}
	started := time.Now()
	v.progress("Updating dataset (one merge for this scan)...")
	if err := v.jsonExport.Flush(dataset); err != nil {
		return reports, err
	}
	for _, entry := range v.retireEntries {
		if referenced[entry.ConversationID] {
			continue
		}
		for path := range entry.Exports {
			if !filepath.IsLocal(path) {
				return reports, fmt.Errorf("unsafe cached export path: %s", path)
			}
			if err := os.Remove(filepath.Join(v.BaseDir, path)); err != nil && !os.IsNotExist(err) {
				return reports, fmt.Errorf("retire old export %s: %w", path, err)
			}
		}
	}
	v.progress(fmt.Sprintf("Dataset complete (%s).", time.Since(started).Round(time.Millisecond)))
	state.Dataset, _ = stamp(dataset)
	data, err = json.Marshal(state)
	if err != nil {
		return reports, err
	}
	if err := exporter.WriteIfChanged(cachePath, data); err != nil {
		return reports, err
	}
	return reports, nil
}

func (v *Vault) unchanged(entry scanEntry, revision string) bool {
	if v.ScanOptions.Full || revision == "" || entry.Revision != revision {
		return false
	}
	for path, previous := range entry.Exports {
		// Cached paths are relative export paths, never source paths.
		if !filepath.IsLocal(path) {
			return false
		}
		current, err := stamp(filepath.Join(v.BaseDir, path))
		if err != nil || current != previous {
			return false
		}
	}
	return true
}

type scanJob struct {
	key, revision, path string
	conv                *models.Conversation
	previous            scanEntry
}
type scanResult struct {
	job      scanJob
	conv     *models.Conversation
	warnings []string
	err      error
}

func (v *Vault) scanSource(source ScanSource, state *scanState) ToolScanReport {
	started := time.Now()
	r := ToolScanReport{ToolName: source.Tool, Location: source.Path}
	v.progress("Scanning " + source.Tool + "...")
	var jobs []scanJob
	if source.Tool == "opencode" {
		revisions := make(map[string]string)
		convs, err := v.opencodeExtractor.ExtractIncremental(source.Path, func(id, revision string) bool {
			r.DiscoveredCount++
			key := sourceKey(source.Tool, source.Path, id)
			revisions[id] = revision
			if v.unchanged(state.Entries[key], revision) {
				r.UnchangedCount++
				return false
			}
			return true
		})
		if err != nil {
			r.Warnings = append(r.Warnings, err.Error())
			r.FailedCount++
		}
		for _, conv := range convs {
			key := sourceKey(source.Tool, source.Path, conv.ID)
			jobs = append(jobs, scanJob{key: key, revision: revisions[conv.ID], conv: conv, previous: state.Entries[key]})
		}
	} else {
		files, err := discoverFiles(source)
		if err != nil {
			r.Warnings = append(r.Warnings, err.Error())
			r.FailedCount++
		}
		for _, path := range files {
			r.DiscoveredCount++
			info, err := stamp(path)
			if err != nil {
				r.Warnings = append(r.Warnings, err.Error())
				r.FailedCount++
				continue
			}
			revision := fmt.Sprintf("%d:%d", info.Size, info.Modified)
			key := sourceKey(source.Tool, path, "")
			entry := state.Entries[key]
			if v.unchanged(entry, revision) {
				r.UnchangedCount++
				continue
			}
			jobs = append(jobs, scanJob{key: key, revision: revision, path: path, previous: entry})
		}
	}
	v.progress(fmt.Sprintf("%s: %d discovered, %d unchanged, %d to process", source.Tool, r.DiscoveredCount, r.UnchangedCount, len(jobs)))
	workers := v.ScanOptions.Workers
	if workers <= 0 {
		workers = min(4, runtime.GOMAXPROCS(0))
	}
	workers = min(workers, len(jobs))
	queue := make(chan scanJob)
	results := make(chan scanResult)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range queue {
				var result scanResult
				result.job = job
				if job.conv != nil {
					result.conv, result.warnings, result.err = v.StoreConversation(job.conv)
				} else {
					raw, err := v.ExtractConversation(job.path, source.Tool)
					if err == nil {
						info, statErr := stamp(job.path)
						if statErr != nil || fmt.Sprintf("%d:%d", info.Size, info.Modified) != job.revision {
							err = fmt.Errorf("source changed during scan; will retry: %s", job.path)
						}
					}
					if err != nil {
						result.err = err
					} else {
						result.conv, result.warnings, result.err = v.StoreConversation(raw)
					}
				}
				results <- result
			}
		}()
	}
	go func() {
		for _, job := range jobs {
			queue <- job
		}
		close(queue)
		wg.Wait()
		close(results)
	}()
	done := 0
	for result := range results {
		done++
		if result.err != nil {
			r.FailedCount++
			r.Warnings = append(r.Warnings, result.err.Error())
			if result.job.previous.ConversationID != "" || result.job.previous.Revision != "" {
				state.Entries[result.job.key] = result.job.previous
			} else {
				delete(state.Entries, result.job.key)
			}
			continue
		}
		entry := scanEntry{Revision: result.job.revision, Exports: make(map[string]fileStamp)}
		if result.conv == nil {
			r.SkippedCount++
		} else {
			r.ImportedCount++
			entry.ConversationID = result.conv.ID
			if result.job.previous.ConversationID != "" || result.job.previous.Revision != "" {
				r.ChangedCount++
			} else {
				r.NewCount++
			}
			v.storeMu.Lock()
			slug := v.exportNames[result.conv.ID].Slug
			v.storeMu.Unlock()
			for _, path := range []string{filepath.Join("conversations", "markdown", slug+".md"), filepath.Join("conversations", "sharegpt", slug+".json")} {
				info, err := stamp(filepath.Join(v.BaseDir, path))
				if err != nil {
					result.err = err
					break
				}
				entry.Exports[path] = info
			}
		}
		r.Warnings = append(r.Warnings, result.warnings...)
		if result.err == nil {
			state.Entries[result.job.key] = entry
			old := result.job.previous
			if old.ConversationID != "" && old.ConversationID != entry.ConversationID {
				v.retireEntries = append(v.retireEntries, old)
			}
		} else {
			r.FailedCount++
			r.Warnings = append(r.Warnings, result.err.Error())
			if result.job.previous.ConversationID != "" || result.job.previous.Revision != "" {
				state.Entries[result.job.key] = result.job.previous
			} else {
				delete(state.Entries, result.job.key)
			}
		}
		if done%25 == 0 || done == len(jobs) {
			v.progress(fmt.Sprintf("%s: processed %d/%d", source.Tool, done, len(jobs)))
		}
	}
	r.Duration = time.Since(started)
	v.progress(fmt.Sprintf("%s finished in %s", source.Tool, r.Duration.Round(time.Millisecond)))
	return r
}

func discoverFiles(source ScanSource) ([]string, error) {
	if source.Tool == "aider" {
		return []string{source.Path}, nil
	}
	var files []string
	if source.Tool == "antigravity" {
		entries, err := os.ReadDir(source.Path)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			for _, name := range []string{"transcript_full.jsonl", "transcript.jsonl"} {
				path := filepath.Join(source.Path, entry.Name(), ".system_generated", "logs", name)
				if _, err := os.Stat(path); err == nil {
					files = append(files, path)
					break
				} else if !os.IsNotExist(err) {
					return files, err
				}
			}
		}
		return files, nil
	}
	if source.Tool != "codex" {
		return nil, fmt.Errorf("unknown scan source %q", source.Tool)
	}
	err := filepath.WalkDir(source.Path, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		name := strings.ToLower(entry.Name())
		if !entry.IsDir() && strings.HasPrefix(name, "rollout-") && strings.HasSuffix(name, ".jsonl") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
