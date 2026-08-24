# Dataset Governance

Only sanitized, problem-solving conversations belong in this repository. Do not commit raw assistant databases, source logs, credentials, personal data, private network details, or proprietary code unless publication is explicitly authorized.

## Review Before Publishing

Run `vault audit` and inspect the generated Markdown and ShareGPT files before every commit. Redaction is defense in depth, not a guarantee that every sensitive value is recognized.

## Retention And Deletion

Retain a conversation only while it remains useful to the shared dataset and publication is authorized. To remove a conversation, delete its matching Markdown and ShareGPT exports and remove the JSONL record with the same `id`, then run `vault audit` before committing the deletion.

## Schema And Attribution

Normalized JSONL records include `schema_version` (currently `1`) and `source_tool`. The stable `id` is derived from sanitized content; it supports repeatable imports without retaining source-system identifiers.
