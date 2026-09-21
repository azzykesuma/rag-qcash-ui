package vault

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"llm-context-vault/pkg/exporter"
	"llm-context-vault/pkg/models"
)

// Bump when extraction, filtering, or export behavior changes.
const scanVersion = "3"

type ScanOptions struct {
	Full                 bool
	Workers              int
	MaxItems             int
	MaxItemsPerTool      int
	Daily                bool
	AuthoritativeSources bool
	Progress             func(string)
}

var ErrDailyScanSkipped = errors.New("daily scan already completed")

type ScanSource struct{ Tool, Path string }

type fileStamp struct {
	Size     int64 `json:"size"`
	Modified int64 `json:"modified"`
}

type scanEntry struct {
	Revision       string               `json:"revision"`
	ConversationID string               `json:"conversation_id,omitempty"`
	Exports        map[string]fileStamp `json:"exports,omitempty"`
	RetryAfterRun  uint64               `json:"retry_after_run,omitempty"`
}

type scanState struct {
	Version          string               `json:"version"`
	Dataset          fileStamp            `json:"dataset"`
	DatasetHash      string               `json:"dataset_hash,omitempty"`
	DailyCompletions map[string]string    `json:"daily_completions,omitempty"`
	Run              uint64               `json:"run,omitempty"`
	Cursors          map[string]string    `json:"cursors,omitempty"`
	SourceCursor     string               `json:"source_cursor,omitempty"`
	Entries          map[string]scanEntry `json:"entries"`
}

func stamp(path string) (fileStamp, error) {
	info, err := os.Stat(path)
	if err != nil {
		return fileStamp{}, err
	}
	return fileStamp{Size: info.Size(), Modified: info.ModTime().UnixNano()}, nil
}

func fileHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func dailyScanKey(version string, sources []ScanSource, options ScanOptions) (string, error) {
	keys := make([]string, 0, len(sources))
	for _, source := range sources {
		path, err := filepath.Abs(source.Path)
		if err != nil {
			return "", err
		}
		keys = append(keys, source.Tool+"\x00"+filepath.Clean(path))
	}
	sort.Strings(keys)
	payload, err := json.Marshal(struct {
		Version      string
		Sources      []string
		Full         bool
		Limit        int
		LimitPerTool int
	}{version, keys, options.Full, options.MaxItems, options.MaxItemsPerTool})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(payload)), nil
}

func rotateStrings(items []string, cursor string) []string {
	if cursor == "" {
		return items
	}
	for i, item := range items {
		if item == cursor {
			return append(append([]string(nil), items[i+1:]...), items[:i+1]...)
		}
	}
	return items
}

func rotateSources(sources []ScanSource, cursor string) []ScanSource {
	if cursor == "" {
		return sources
	}
	for i, source := range sources {
		if sourceKey(source.Tool, source.Path, "") == cursor {
			return append(append([]ScanSource(nil), sources[i+1:]...), sources[:i+1]...)
		}
	}
	return sources
}

func sourceKey(tool, path, session string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(tool+"\x00"+path+"\x00"+session)))
}

func parseLockPID(path string) int {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "pid=") {
			var pid int
			if _, err := fmt.Sscanf(line, "pid=%d", &pid); err == nil && pid > 0 {
				return pid
			}
		}
	}
	return 0
}

func (v *Vault) lockWriter() (func(), error) {
	dir := filepath.Join(v.BaseDir, ".vault")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, err
	}
	path := filepath.Join(dir, "write.lock")

	for attempts := 0; attempts < 2; attempts++ {
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if err == nil {
			fmt.Fprintf(f, "pid=%d\n", os.Getpid())
			f.Close()
			return func() { _ = os.Remove(path) }, nil
		}

		if !os.IsExist(err) {
			return nil, fmt.Errorf("vault writer lock %s: %w", path, err)
		}

		pid := parseLockPID(path)
		if pid > 0 && !isProcessAlive(pid) {
			v.progress(fmt.Sprintf("Cleaning up stale writer lock from terminated PID %d...", pid))
			_ = os.Remove(path)
			continue
		}

		return nil, fmt.Errorf("vault writer lock %s: locked by active PID %d; if a previous process crashed, remove the lock only after confirming it is no longer running", path, pid)
	}

	return nil, fmt.Errorf("vault writer lock %s: failed to acquire lock", path)
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
	if v.ScanOptions.Daily && (v.ScanOptions.Full || v.ScanOptions.MaxItems <= 0 && v.ScanOptions.MaxItemsPerTool <= 0) {
		return nil, fmt.Errorf("daily scans require a bounded item limit")
	}
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
	dailyKey := ""
	if v.ScanOptions.Daily {
		dailyKey, err = dailyScanKey(version, sources, v.ScanOptions)
		if err != nil {
			return nil, err
		}
		if state.DailyCompletions[dailyKey] == time.Now().Format("2006-01-02") {
			return nil, ErrDailyScanSkipped
		}
	}
	state.Run++
	currentDataset, datasetErr := stamp(dataset)
	invalidated := state.Version != version || state.Entries == nil
	rebuildDataset := false
	if !invalidated && len(state.Entries) > 0 {
		switch {
		case os.IsNotExist(datasetErr):
			v.progress("Dataset is missing; rebuilding from source sessions.")
			invalidated = true
			rebuildDataset = true
		case datasetErr != nil:
			return nil, datasetErr
		case state.Dataset != currentDataset && state.DatasetHash != "":
			currentHash, err := fileHash(dataset)
			if err != nil {
				return nil, err
			}
			if currentHash != state.DatasetHash {
				v.progress("Dataset content changed outside the scanner; rebuilding from source sessions.")
				invalidated = true
				rebuildDataset = true
			} else {
				// File timestamps can change after Git operations or file restoration.
				// Matching content must not force every source through extraction again.
				state.Dataset = currentDataset
			}
		}
	}
	if state.Entries == nil {
		state.Entries = make(map[string]scanEntry)
	}
	if state.Cursors == nil {
		state.Cursors = make(map[string]string)
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
	if rebuildDataset {
		if !v.ScanOptions.AuthoritativeSources {
			return nil, fmt.Errorf("dataset recovery requires a unified scan; run `vault scan`")
		}
		v.jsonExport.BeginRebuild()
	} else {
		v.jsonExport.BeginBatch()
	}
	var reports []ToolScanReport
	seenSources := make(map[string]bool)
	var prepared []ScanSource
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
		prepared = append(prepared, source)
	}
	bounded := !v.ScanOptions.Full && !rebuildDataset
	if bounded && (v.ScanOptions.MaxItems > 0 || v.ScanOptions.MaxItemsPerTool > 0) {
		prepared = rotateSources(prepared, state.SourceCursor)
	}
	globalBudget := scanBudget{remaining: v.ScanOptions.MaxItems, limited: bounded && v.ScanOptions.MaxItems > 0}
	toolBudgets := make(map[string]*scanBudget)
	for _, source := range prepared {
		key := sourceKey(source.Tool, source.Path, "")
		budget := &globalBudget
		if bounded && v.ScanOptions.MaxItemsPerTool > 0 {
			budget = toolBudgets[source.Tool]
			if budget == nil {
				budget = &scanBudget{remaining: v.ScanOptions.MaxItemsPerTool, limited: true}
				toolBudgets[source.Tool] = budget
			}
		}
		report := v.scanSource(source, &state, budget, rebuildDataset)
		reports = append(reports, report)
		if bounded && report.ImportedCount+report.SkippedCount+report.FailedCount > 0 {
			state.SourceCursor = key
		}
	}
	if rebuildDataset {
		if len(prepared) == 0 {
			return reports, fmt.Errorf("dataset rebuild requires at least one available source")
		}
		for _, report := range reports {
			if report.FailedCount > 0 || report.DeferredCount > 0 {
				return reports, fmt.Errorf("dataset rebuild failed; previous dataset was preserved")
			}
		}
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
	updatedDataset, updatedErr := stamp(dataset)
	if updatedErr == nil {
		if state.DatasetHash == "" || updatedDataset != currentDataset {
			state.DatasetHash, err = fileHash(dataset)
			if err != nil {
				return reports, err
			}
		}
		state.Dataset = updatedDataset
	} else if !os.IsNotExist(updatedErr) {
		return reports, updatedErr
	}
	failed := false
	for _, report := range reports {
		failed = failed || report.FailedCount > 0
	}
	if v.ScanOptions.Daily && !failed {
		if state.DailyCompletions == nil {
			state.DailyCompletions = make(map[string]string)
		}
		state.DailyCompletions[dailyKey] = time.Now().Format("2006-01-02")
	}
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

type scanBudget struct {
	remaining int
	limited   bool
}

func (b *scanBudget) take() bool {
	if !b.limited {
		return true
	}
	if b.remaining <= 0 {
		return false
	}
	b.remaining--
	return true
}

func (v *Vault) scanSource(source ScanSource, state *scanState, budget *scanBudget, ignoreRetry bool) ToolScanReport {
	started := time.Now()
	r := ToolScanReport{ToolName: source.Tool, Location: source.Path}
	v.progress("Scanning " + source.Tool + "...")
	var jobs []scanJob
	cursorKey := sourceKey(source.Tool, source.Path, "")
	lastSelected := ""
	if source.Tool == "opencode" {
		revisions := make(map[string]string)
		convs, err := v.opencodeExtractor.ExtractIncrementalFrom(source.Path, state.Cursors[cursorKey], func(id, revision string) bool {
			r.DiscoveredCount++
			key := sourceKey(source.Tool, source.Path, id)
			revisions[id] = revision
			entry := state.Entries[key]
			if v.unchanged(entry, revision) {
				r.UnchangedCount++
				return false
			}
			if !ignoreRetry && entry.RetryAfterRun >= state.Run {
				r.DeferredCount++
				return false
			}
			if !budget.take() {
				r.DeferredCount++
				return false
			}
			lastSelected = id
			return true
		})
		if lastSelected != "" {
			state.Cursors[cursorKey] = lastSelected
		}
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
		files = rotateStrings(files, state.Cursors[cursorKey])
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
			if !ignoreRetry && entry.RetryAfterRun >= state.Run {
				r.DeferredCount++
				continue
			}
			if !budget.take() {
				r.DeferredCount++
				continue
			}
			lastSelected = path
			jobs = append(jobs, scanJob{key: key, revision: revision, path: path, previous: entry})
		}
		if lastSelected != "" {
			state.Cursors[cursorKey] = lastSelected
		}
	}
	v.progress(fmt.Sprintf("%s: %d discovered, %d unchanged, %d deferred, %d to process", source.Tool, r.DiscoveredCount, r.UnchangedCount, r.DeferredCount, len(jobs)))
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
			entry := result.job.previous
			if budget.limited {
				entry.RetryAfterRun = state.Run + 1
			}
			state.Entries[result.job.key] = entry
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
			entry := result.job.previous
			if budget.limited {
				entry.RetryAfterRun = state.Run + 1
			}
			state.Entries[result.job.key] = entry
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
