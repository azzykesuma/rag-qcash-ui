# Add pre-commit workflow to skill.md

- **Author**: `azzy`
- **ID**: `ses_0356bad83ffec996XWy9i8Ciqj`
- **Source Tool**: `opencode`
- **Date**: `2026-08-04 09:22:25`
- **Tags**: `coding`, `assistant`, `opencode`

---

## Turn 1: User

---
name: pre-commit
description: Execute pre-commit checks: ask for the Jira ticket number, run yarn lint, run yarn build, inspect uncommitted code for potential Sonar issues ("I'm about to commit, will this create sonar issues?"), and perform a quality gateway check for new code. Triggered when preparing for a git commit or running pre-commit checks.
---

# Pre-Commit Check Skill

This skill guides the agent through a comprehensive pre-commit routine before committing code changes to Git.

## Workflow & Steps

### Step 1: Prompt for Jira Ticket Number
- Ask the user for the relevant Jira ticket number (e.g., `QCASH-1234` or `JIRA-5678`).
- If the ticket number is already provided in the prompt or conversation history, confirm and use it.
- Store the Jira ticket number to reference in the pre-commit report and commit message format.

### Step 2: Run Code Quality Checks & Build
Execute the project verification scripts:
1. **Lint Check**: Run `yarn lint` to check for ESLint, style, or formatting errors.
2. **Build Check**: Run `yarn build` to verify Next.js/TypeScript compilation.

If any command fails, summarize the error log, stop the pre-commit process, and request fix before proceeding.

### Step 3: Sonar Issue Check & Quality Gateway Scan for New Code
Perform a scan/check on newly added or modified code (`git status`, `git diff`):
1. **Check Current Code for Sonar Issues**: Inspect staged and uncommitted code to answer: *"I'm about to commit, will this create sonar issues?"*
2. **Evaluate Quality Gate Conditions for New Code**:
   - **Bugs & Vulnerabilities**: Check for potential null dereferences, unhandled exceptions, or security risks. (Threshold: 0 Blocker/Critical)
   - **Code Smells**: Check for unused variables, console statements, or dead code. (Threshold: 0 High/Critical)
   - **Duplicated Lines**: Ensure no duplicate logic added. (Threshold: < 3%)
   - **Security Hotspots**: 0 unreviewed hotspots.
3. **Generate Pre-Commit & Sonar Verification Report**: Present a structured summary in Markdown:

```markdown
### 🛡️ Pre-Commit Verification Report

- **Jira Ticket**: [TICKET-ID]
- **Yarn Lint**: ✅ Passed / ❌ Failed
- **Yarn Build**: ✅ Passed / ❌ Failed
- **Sonar Issues Check**: 🟢 NO NEW ISSUES / 🔴 NEW ISSUES DETECTED
- **Sonar Quality Gate (New Code)**: 🟢 PASSED / 🔴 FAILED

#### Sonar Quality Gate Breakdown
| Metric | Threshold | Scan Result | Status |
| :--- | :--- | :--- | :--- |
| **New Bugs** | 0 | 0 | 🟢 OK |
| **Vulnerabilities** | 0 | 0 | 🟢 OK |
| **Code Smells** | 0 | 0 | 🟢 OK |
| **Duplications** | < 3% | 0.0% | 🟢 OK |
| **Security Hotspots** | 0 | 0 | 🟢 OK |
```

4. If all checks pass, notify the user that the codebase is ready for commit. add this into your skill.md

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **skill**: Tool: skill

</details>

I’ll update the existing `pre-commit` skill definition, preserving its current structure while incorporating the expanded Sonar inspection and report requirements.

---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>

The existing skill already contains the base workflow. I’m replacing its mock-only Sonar step with an explicit uncommitted-code issue review, adding the separate issue status to the report, and updating the trigger description.

---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 6: Assistant

Updated `~\.config\opencode\skills\pre-commit\SKILL.md` with the expanded Sonar issue and quality-gate checks.

Restart OpenCode for the updated skill to take effect.

---

## Turn 7: User

{
  "mcpServers": {
    "context7": {
      "serverUrl": "https://mcp.context7.com/mcp",
      "headers": {
        "CONTEXT7_API_KEY": "[ATLASSIAN_TOKEN_REDACTED]"
      }
    },
    "playwright": {
      "command": "npx.cmd",
      "args": [
        "@playwright/mcp@latest"
      ],
      "env": {
        "npm_config_registry": "https://registry.npmjs.org/"
      }
    },
    "sonarqube": {
      "command": "node",
      "args": [
        "C:\\Users\\90181271\\.gemini\\antigravity-cli\\scratch\\node_modules\\sonarqube-mcp-server\\dist\\index.js"
      ],
      "env": {
        "NODE_TLS_REJECT_UNAUTHORIZED": "0",
        "SONARQUBE_ORG": "apps.sonarqube.user-vendor",
        "SONARQUBE_TOKEN": "[ATLASSIAN_TOKEN_REDACTED]",
        "SONARQUBE_URL": "https://internal-service.example.com"
      }
    },
    "atlassian-bitbucket-dc": {
      "command": "node",
      "args": [
        "C:\\Users\\90181271\\AppData\\Local\\npm-cache\\_npx\\8ed5ea13b8a1daeb\\node_modules\\@atlassian-dc-mcp\\bitbucket\\build\\index.js"
      ],
      "env": {
        "BITBUCKET_HOST": "https://internal-service.example.com",
        "BITBUCKET_API_TOKEN": "[ATLASSIAN_TOKEN_REDACTED]",
        "NODE_TLS_REJECT_UNAUTHORIZED": "0"
      }
    },
    "atlassian-jira-dc": {
      "command": "npx",
      "args": [
        "-y",
        "@atlassian-dc-mcp/jira"
      ],
      "env": {
        "BITBUCKET_HOST": "https://internal-service.example.com/",
        "BITBUCKET_API_TOKEN": "[ATLASSIAN_TOKEN_REDACTED]",
        "NODE_TLS_REJECT_UNAUTHORIZED": "0"
      }
    }
  }
}  add these into mcp config as well

---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (2)</summary>

- **glob**: Tool: glob
- **glob**: Tool: glob

</details>

I’ll translate these server definitions into OpenCode’s `mcp` schema, merge them into the existing global config, and preserve unrelated settings. Local server commands will become command arrays, while Context7 will use the remote server shape.

---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **read**: Tool: read

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **apply_patch**: Tool: apply_patch

</details>

The global config currently contains only the schema declaration, so there are no existing MCP entries to reconcile. I’m adding all five servers now using OpenCode’s required `type`, `url`, `command`, and `environment` fields.

---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **bash**: Tool: bash

</details>

The config has been written. I’m validating it through OpenCode itself so schema or startup errors are caught before you restart; this may also reveal server-specific connection failures without changing the configuration.

---

## Turn 12: Assistant

Added all five MCP servers to `~\.config\opencode\opencode.jsonc`.

`opencode mcp list` confirms all five connect successfully:

- Context7
- Playwright
- SonarQube
- Atlassian Bitbucket DC
- Atlassian Jira DC

Restart the current OpenCode session to load them. Credentials are stored as plaintext in the config.

---

