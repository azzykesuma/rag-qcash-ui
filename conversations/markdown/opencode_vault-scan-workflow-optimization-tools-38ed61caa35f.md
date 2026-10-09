# Vault scan workflow optimization tools

- **ID**: `opencode-212b0c67c82b87515cbf44cacf9e936a4350f8069afc1f79c59238ed61caa35f`
- **Source Tool**: `opencode`
- **Date**: `2026-10-01 11:27:32`
- **Project**: `qcash-ui`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

check all of my vault scan, and recommend me the tools that i can create with you to optimize my workflow

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>

I’ll inventory the vault’s organization and conversation history, then turn recurring work patterns into concrete, buildable tool recommendations.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (4)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob
- **glob**: Tool: glob

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **task**: Tool: task

</details>

The vault contains substantially more than the first directory listing can display, with recurring `qcash-ui`, Git, Bamboo, SonarQube, API, and environment work visible already. I’m refreshing the vault scan and analyzing the full exported corpus by themes rather than relying on filenames alone.

---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **todowrite**: Tool: todowrite

</details>



---

## Turn 8: Assistant

Scanned the complete vault: 250 Antigravity, 205 OpenCode, and 5 Codex sessions. The refresh imported 3 new/changed OpenCode exports and passed the full privacy audit.

Your highest-leverage tools, based on recurring work:

1. **MFE Release Train Planner**
   - Maps affected repositories, release branches, prerequisite commits, PR order, deployment order, and rollback points.
   - Best fit for repeated multi-repo promotion and cherry-pick work.

2. **Federation Compatibility Linter**
   - Checks all QCash MFEs for shared dependency/version drift, singleton settings, federation runtime versions, remotes, and dynamic import ordering.
   - Prevents your recurring auth-context and production-only federation failures.

3. **QShip 2: Resumable Delivery Orchestrator**
   - Evolve `qship` into a stateful workflow: branch validation, change/test/Sonar evidence, required reviewers, PR creation, promotion PRs, and handoff message.
   - Eliminates repeated delivery checklists and safely resumes interrupted releases.

4. **Cross-Repo Pattern Finder**
   - Searches local QCash repos, Bitbucket, and vault decisions together for proven implementations.
   - Return implementation source, branch/commit, and a concise comparison.
   - Targets the frequent “how does another repo handle this?” investigations.

5. **Bamboo Failure Correlator**
   - Connects failed logs to the likely Dockerfile, package script, Bamboo configuration, changed commit, root cause, and verification command.
   - Build it on top of the existing Build Doctor integration.

6. **Sonar New-Code Test Planner**
   - Reads new-code coverage issues, maps them to exact source paths, finds nearby test conventions, and creates only the smallest tests needed for the quality gate.

7. **Production MFE Artifact Inspector**
   - Inspects deployed `remoteEntry.js` and chunks, maps them to release refs/configuration, and detects stale or cross-environment remotes.
   - Strong fit for hard-to-reproduce production MFE incidents.

8. **QA Flow Extractor with Playwright**
   - Derives a route/form/state QA matrix from changed code, runs prioritized desktop/mobile smoke tests, and saves screenshots, console logs, and network failures.

9. **API Contract and Error Flow Mapper**
   - Traces an endpoint through service layer, payload, Axios/interceptors, UI state, process IDs, error dialogs, and tests.
   - Useful for recurring API/error-handling investigations.

10. **Vault Decision Digest**
   - Produces source-cited project digests: established decisions, incident playbooks, recurring failures, and newly detected automation opportunities.
   - Turns the vault into active institutional memory rather than archive-only search.

Start with **MFE Release Train Planner**. It has the clearest ROI, uses your existing Bitbucket/Bamboo/Sonar integrations, and can later become a first-class phase of `QShip 2`.

Evidence in the vault strongly centers on repository investigation, frontend fixes, delivery automation, MFE debugging, CI failures, and Sonar remediation.

---

