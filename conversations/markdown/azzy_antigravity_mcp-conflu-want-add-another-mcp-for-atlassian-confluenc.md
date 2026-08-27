# mcp conflu

- **Author**: `azzy`
- **ID**: `e6803251-82c4-4eba-9035-537b893021b9`
- **Source Tool**: `antigravity`
- **Date**: `2026-08-27 14:24:48`
- **Tags**: `coding`, `assistant`, `antigravity`

---

## Turn 1: User

mcp conflu
i want to add another mcp for atlassian confluence , you can refer to the other atlassian mcp
for the token, i want you to generate a safe script for me to add it

---

## Turn 2: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 3: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 4: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 5: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 6: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 7: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 8: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 9: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 10: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>

I'm checking the `@atlassian-dc-mcp/confluence` package details and the existing Atlassian MCP configuration to prepare a safe script for your Confluence MCP setup.

---

## Turn 11: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 12: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 13: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 14: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 15: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **list_dir**: 

</details>



---

## Turn 16: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 17: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 18: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **find_by_name**: 

</details>



---

## Turn 19: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 20: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **list_dir**: 

</details>



---

## Turn 21: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **list_dir**: 

</details>



---

## Turn 22: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **list_dir**: 

</details>



---

## Turn 23: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **write_to_file**: 

</details>



---

## Turn 24: Assistant

To configure the **Atlassian Confluence DC MCP** server matching your existing Jira and Bitbucket setups, here is the setup structure and a safe configuration script.

---

### Configuration Pattern (Aligned with Jira & Bitbucket)

The MCP server uses [`@atlassian-dc-mcp/confluence`](https://github.com/b1ff/atlassian-dc-mcp) with the following structure:

```json
{
  "atlassian-confluence-dc": {
    "command": "npx",
    "args": [
      "-y",
      "@atlassian-dc-mcp/confluence"
    ],
    "env": {
      "CONFLUENCE_HOST": "https://internal-service.example.com",
      "CONFLUENCE_API_TOKEN": "<YOUR_TOKEN>",
      "CONFLUENCE_TOKEN": "<YOUR_TOKEN>",
      "NODE_TLS_REJECT_UNAUTHORIZED": "0"
    }
  }
}
```

---

### Safe Token Setup Script

To keep your token **private and secure** (without pasting it into LLM chat or exposing it in shell history), a script has been saved at:
[`setup-confluence-mcp.ps1`](file:///C:~/.gemini/antigravity-cli/brain/e6803251-82c4-4eba-9035-537b893021b9/scratch/setup-confluence-mcp.ps1)

#### What this script does:
1. Prompts for your Personal Access Token with **masked input** (`-AsSecureString`) so it won't appear on screen or in your PowerShell command history.
2. Backs up existing `mcp_config.json` files (`.bak`).
3. Automatically updates both `.gemini\antigravity-cli\mcp_config.json` and `.gemini\config\mcp_config.json`.
4. Creates `%USERPROFILE%\.atlassian-dc-mcp\confluence.env` as a local fallback.
5. Immediately flushes memory variables after saving.

---

### How to Run

Open PowerShell in your terminal and execute:

```powershell
& "~\.gemini\antigravity-cli\brain\e6803251-82c4-4eba-9035-537b893021b9\scratch\setup-confluence-mcp.ps1"
```

*Or copy and paste the script directly into a PowerShell window:*

```powershell
# Interactive Confluence MCP Setup
$HostUrl = "https://internal-service.example.com"
$enteredHost = Read-Host "Confluence Host URL [Default: $HostUrl]"
if (-not [string]::IsNullOrWhiteSpace($enteredHost)) { $HostUrl = $enteredHost.Trim().TrimEnd('/') }

$secureToken = Read-Host -Prompt "Enter your Confluence Personal Access Token (PAT)" -AsSecureString
if ($null -eq $secureToken -or $secureToken.Length -eq 0) { Write-Error "Token cannot be empty."; return }

$bstr = [System.Runtime.InteropServices.Marshal]::SecureStringToBSTR($secureToken)
$plainToken = [System.Runtime.InteropServices.Marshal]::PtrToStringAuto($bstr)
[System.Runtime.InteropServices.Marshal]::ZeroFreeBSTR($bstr)

$confluenceConfig = [ordered]@{
    command = "npx"
    args    = @("-y", "@atlassian-dc-mcp/confluence")
    env     = [ordered]@{
        CONFLUENCE_HOST              = $HostUrl
        CONFLUENCE_API_TOKEN         = $plainToken
        CONFLUENCE_TOKEN             = $plainToken
        NODE_TLS_REJECT_UNAUTHORIZED = "0"
    }
}

$targetConfigs = @(
    (Join-Path $env:USERPROFILE ".gemini\antigravity-cli\mcp_config.json"),
    (Join-Path $env:USERPROFILE ".gemini\config\mcp_config.json")
)

foreach ($cfgPath in $targetConfigs) {
    if (Test-Path $cfgPath) {
        Copy-Item -Path $cfgPath -Destination "$cfgPath.bak" -Force
        $data = Get-Content -Path $cfgPath -Raw -Encoding UTF8 | ConvertFrom-Json -AsHashtable
        if (-not $data.ContainsKey("mcpServers")) { $data["mcpServers"] = [ordered]@{} }
        $data["mcpServers"]["atlassian-confluence-dc"] = $confluenceConfig
        [System.IO.File]::WriteAllText($cfgPath, ($data | ConvertTo-Json -Depth 10), [System.Text.Encoding]::UTF8)
        Write-Host "[✓] Updated: $cfgPath" -ForegroundColor Green
    }
}

$dcMcpDir = Join-Path $env:USERPROFILE ".atlassian-dc-mcp"
if (-not (Test-Path $dcMcpDir)) { New-Item -ItemType Directory -Path $dcMcpDir -Force | Out-Null }
$envFilePath = Join-Path $dcMcpDir "confluence.env"
[System.IO.File]::WriteAllText($envFilePath, "CONFLUENCE_HOST=$HostUrl`nCONFLUENCE_API_TOKEN=$plainToken`nCONFLUENCE_TOKEN=$plainToken`nNODE_TLS_REJECT_UNAUTHORIZED=0", [System.Text.Encoding]::UTF8)
Write-Host "[✓] Created fallback env: $envFilePath" -ForegroundColor Green

$plainToken = $null
$secureToken = $null
[System.GC]::Collect()
Write-Host "`n✓ Confluence MCP setup complete! Restart your session to activate tools." -ForegroundColor Cyan
```

---

### Available Tools After Setup

Once configured and restarted, the following Confluence MCP tools will be available:
- `confluence_searchContent`: Search pages, blog posts, and attachments with CQL queries
- `confluence_getContent`: Retrieve content by ID with body formatting
- `confluence_createContent`: Create pages or blog posts
- `confluence_updateContent`: Edit existing Confluence content
- `confluence_searchSpace`: List and search spaces
- `confluence_downloadAttachment`: Retrieve Confluence page attachments

---

## Turn 25: User

can you just make a simple bash script ?

---

## Turn 26: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **write_to_file**: 

</details>



---

## Turn 27: Assistant

Here is a simple bash script. It prompts for your token silently (`read -s`) without echoing it to the terminal, creates automatic backups, and writes the configuration.

Saved at: [`setup-confluence-mcp.sh`](file:///C:~/.gemini/antigravity-cli/brain/e6803251-82c4-4eba-9035-537b893021b9/scratch/setup-confluence-mcp.sh)

---

### Bash Script

```bash
#!/usr/bin/env bash
set -e

echo "=== Atlassian Confluence DC MCP Setup ==="

# 1. Prompt for Host URL (defaults to https://internal-service.example.com)
read -r -p "Confluence Host URL [https://internal-service.example.com]: " HOST_INPUT
CONFLUENCE_HOST="${HOST_INPUT:-https://internal-service.example.com}"
CONFLUENCE_HOST="${CONFLUENCE_HOST%/}"

# 2. Prompt for Token silently (hidden input)
read -s -r -p "Enter Confluence Personal Access Token (PAT): " CONFLUENCE_TOKEN
echo ""

if [ -z "$CONFLUENCE_TOKEN" ]; then
    echo "Error: Token cannot be empty."
    exit 1
fi

# Target config files
TARGET_CONFIGS=(
    "$HOME/.gemini/antigravity-cli/mcp_config.json"
    "$HOME/.gemini/config/mcp_config.json"
)

# 3. Update MCP configs safely using node
for CFG in "${TARGET_CONFIGS[@]}"; do
    if [ -f "$CFG" ]; then
        cp "$CFG" "${CFG}.bak"
        node -e '
            const fs = require("fs");
            const file = process.argv[1];
            const host = process.argv[2];
            const token = process.argv[3];
            
            const raw = fs.readFileSync(file, "utf8");
            const data = JSON.parse(raw);
            data.mcpServers = data.mcpServers || {};
            data.mcpServers["atlassian-confluence-dc"] = {
                command: "npx",
                args: ["-y", "@atlassian-dc-mcp/confluence"],
                env: {
                    CONFLUENCE_HOST: host,
                    CONFLUENCE_API_TOKEN: token,
                    CONFLUENCE_TOKEN: token,
                    NODE_TLS_REJECT_UNAUTHORIZED: "0"
                }
            };
            fs.writeFileSync(file, JSON.stringify(data, null, 2) + "\n", "utf8");
        ' "$CFG" "$CONFLUENCE_HOST" "$CONFLUENCE_TOKEN"
        echo "[✓] Updated: $CFG (Backup: ${CFG}.bak)"
    fi
done

# 4. Also write ~/.atlassian-dc-mcp/confluence.env
mkdir -p "$HOME/.atlassian-dc-mcp"
cat <<EOF > "$HOME/.atlassian-dc-mcp/confluence.env"
CONFLUENCE_HOST=$CONFLUENCE_HOST
CONFLUENCE_API_TOKEN=$CONFLUENCE_TOKEN
CONFLUENCE_TOKEN=$CONFLUENCE_TOKEN
NODE_TLS_REJECT_UNAUTHORIZED=0
EOF
chmod 600 "$HOME/.atlassian-dc-mcp/confluence.env" 2>/dev/null || true
echo "[✓] Created fallback env: $HOME/.atlassian-dc-mcp/confluence.env"

# 5. Clear variable from memory
unset CONFLUENCE_TOKEN

echo ""
echo "✓ Setup complete! Restart your IDE/session to load the Confluence MCP tools."
```

---

### How to Run in Git Bash / Terminal

```bash
bash "$HOME/.gemini/antigravity-cli/brain/e6803251-82c4-4eba-9035-537b893021b9/scratch/setup-confluence-mcp.sh"
```

---

## Turn 28: User

check the confluence mcp connection now

---

## Turn 29: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 30: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 31: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **call_mcp_tool**: 

</details>



---

## Turn 32: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 33: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 34: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 35: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **manage_task**: 

</details>



---

## Turn 36: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 37: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **run_command**: 

</details>



---

## Turn 38: Assistant

<details><summary>🔧 Tool Invocations (1)</summary>

- **view_file**: 

</details>



---

## Turn 39: Assistant

The **Atlassian Confluence DC MCP** connection is configured and verified.

### Verification Results

| Check | Result | Details |
| :--- | :--- | :--- |
| **Target Host** | `https://internal-service.example.com` | Reachable |
| **Authentication** | `HTTP 200 OK` | Personal Access Token validated |
| **Config Locations** | Updated & Verified | `~/.gemini/antigravity-cli/mcp_config.json`<br>`~/.gemini/config/mcp_config.json`<br>`~/.atlassian-dc-mcp/confluence.env` |
| **Sample Query** | Success | Retrieved spaces (e.g. `NDS`) & pages (e.g. `Kajian Collaboration Tools`, `Enterprise Service Bus (ESB)`) |

---

### Confluence MCP Tools Available

- `confluence_searchContent` — CQL search across pages, blog posts, and attachments
- `confluence_getContent` — Retrieve content bodies and metadata by ID
- `confluence_createContent` — Create new pages/posts in a space
- `confluence_updateContent` — Update existing pages
- `confluence_searchSpace` — Search and list spaces
- `confluence_downloadAttachment` — Retrieve attachment files

---

