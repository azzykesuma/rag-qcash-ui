# Authenticate with GitHub CLI without writing a personal access token into project or MCP config files.
# Install GitHub CLI from https://cli.github.com/ before running this script.

if (-not (Get-Command gh -ErrorAction SilentlyContinue)) {
    Write-Error "GitHub CLI (gh) is required. Install it from https://cli.github.com/ and run this script again."
    exit 1
}

gh auth login --web --git-protocol https
if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}

gh auth status
