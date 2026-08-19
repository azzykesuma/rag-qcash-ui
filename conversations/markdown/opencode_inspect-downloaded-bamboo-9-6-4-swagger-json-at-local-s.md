# Extract Bamboo REST endpoints (@explore subagent)

- **ID**: `ses_0165ce7f5ffeD6bIO0kaTqPHJW`
- **Source Tool**: `opencode`
- **Date**: `2026-08-10 10:06:47`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

Inspect the downloaded Bamboo 9.6.4 Swagger JSON at ~\.local\share\opencode\tool-output\tool_fe9a297d800146Fy548PcUsIMq. Research only; do not edit files. Return the exact GET endpoint paths, query parameters, response media types, and relevant response structure names for: searching plans, listing latest results for a plan or branch, fetching one result expanded with stages/jobs, and downloading result/job logs. Also note Bearer PAT authentication if represented. Be concise but exact.

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (3)</summary>

- **bash**: Tool: bash
- **bash**: Tool: bash
- **bash**: Tool: bash

</details>



---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 15: Assistant

The spec’s server base is `/rest/`; paths below include that prefix.

### Search plans
`GET /rest/api/latest/search/plans`

Optional query parameters:

- `max-result`: integer
- `searchTerm`: string
- `start-index`: integer
- `permission`: `READ | WRITE | BUILD | CLONE | VIEW_CONFIGURATION | ADMINISTRATION` — default described as `READ`
- `type`: `CHAIN | JOB | CHAIN_BRANCH`
- `fuzzy`: boolean — default described as `false`

`200 Content-Type: application/json`

Structure: `SearchResultsList` → `searchResults[]: SearchResult` → `entity: SearchResultEntity`.

### Latest results

**Plan or job plan**

`GET /rest/api/latest/result/{projectKey}-{buildKey}`

**Named branch**

`GET /rest/api/latest/result/{projectKey}-{buildKey}/branch/{branchName}`

Shared optional query parameters:

- `includeAllStates`: boolean
- `continuable`: boolean
- `issueKey`: comma-separated string
- `max-results`: integer
- `start-index`: integer, zero-based
- `label`: comma-separated string
- `buildstate`: `Unknown | Successful | Failed`
- `favourite`: string
- `lifeCycleState`: comma-separated `Pending | Queued | InProgress | Finished | NotBuilt`
- `expand`: string; documented expansions include `artifacts`, `comments`, `labels`, `jiraIssues`, `stages`, prefixed with `results.result`

`200/403/404 Content-Type: application/json`

Structure: `RestResults` → `results: RestResultList` → `results[]: Result`.

### One result, expanded through stages/jobs
`GET /rest/api/latest/result/{projectKey}-{buildKey}-{buildNumber}`

Optional query parameters:

- `expand`: string
- `buildPositions`: string

For stages and job results:

```text
expand=stages.stage.results.result
```

Other documented expansions: `changes`, `metadata`, `artifacts`, `comments`, `labels`, `jiraIssues`, `stages`, `logEntries`. `logEntries` is job-result-only. `buildNumber` may be `latest`.

`200/403/404 Content-Type: application/json`

Structure: `Result` → `stages: RestStageResultList` → `stages[]: RestStageResult` → `results: RestResultList` → `results[]: Result` (jobs).

Alias also represented:

`GET /rest/api/latest/result/{projectKey}-{buildKey}/{buildNumber}` with optional `expand`; response is `Result`.

### Result/job logs
No dedicated result/job log-download GET path is represented in this Swagger. The represented API mechanism is the single-result endpoint above with:

```text
expand=logEntries
```

for a job result, returning `application/json`. Although documented by the operation, `Result` does not declare a `logEntries` property/schema. Consequently, no `/download/...` path or downloadable-log media type can be established from this JSON.

### Authentication
Bearer/PAT authentication is **not represented**. The only declared security scheme is HTTP Basic (`basic`), and the document has no global or operation-level security requirement.

Source: `~\.local\share\opencode\tool-output\tool_fe9a297d800146Fy548PcUsIMq`

---

