# Repository Instructions

## Verification

- Requires Go 1.25 (`go.mod`). CI runs, in order: `go build -o vault ./cmd/vault`, `go test ./...`, `go vet ./...`, then `./vault audit`.
- Run one package with `go test ./pkg/vault`; run one test with `go test ./pkg/vault -run '^TestName$'`.
- On Windows, build with `go build -o vault.exe ./cmd/vault` and audit with `.\vault.exe audit`.
- Treat the audit text as authoritative: the current `audit` command reports findings but does not exit nonzero. Confirm it prints `Safe to publish!` rather than relying only on the process status.

## Code Map and Coupling

- `cmd/vault` owns CLI parsing and command orchestration. `pkg/extractor` normalizes assistant-specific inputs; `pkg/sanitizer` redacts them; `pkg/vault` coordinates filtering, incremental scans, storage, and search; `pkg/exporter` writes durable output; `pkg/webui` serves the embedded local dashboard.
- Adding an extractor requires more than implementing `Extractor`: register it in `vault.New`, add unified/tool-specific source discovery where applicable, and cover extraction plus scan behavior.
- When extraction, trivial-session filtering, or export semantics change, bump `scanVersion` in `pkg/vault/scan.go`; otherwise cached sources may not be reprocessed. Sanitizer rule/config changes already invalidate scans through `Sanitizer.Fingerprint()`.
- A stored conversation has three coupled exports: `conversations/markdown/*.md`, `conversations/sharegpt/*.json`, and its record in `conversations/dataset.jsonl`. Keep all three consistent; stable IDs are derived from sanitized content. For deletion, remove both files and the matching JSONL record.
- `.vault/` is disposable, ignored local state (scan checkpoints, writer lock, search index, and project aliases). Do not commit it. Do not remove `.vault/write.lock` unless the recorded process is confirmed dead.

## Privacy Boundary

- Never commit raw assistant databases/logs, credentials, personal data, private network details, machine-specific paths, or unauthorized proprietary code. Only sanitized exports belong under `conversations/`.
- Before committing dataset changes, run the full audit and manually inspect the generated Markdown and ShareGPT output; bounded scans deliberately defer the repository-wide audit.
- Preserve atomic/batched export behavior: scans hold one writer lock, merge `dataset.jsonl` once, and checkpoint only after a successful flush. Tests in `pkg/vault/scan_test.go` and `pkg/exporter/jsonl_test.go` cover recovery and must accompany changes to this flow.
