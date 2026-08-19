# Masked GitHub Personal Access Token Configuration Script
# Safely prompts for token, tests authentication against GitHub API,
# and updates Antigravity and OpenCode MCP configs without logging to chat or disk.

$secureToken = Read-Host "GitHub Personal Access Token (PAT)" -AsSecureString
if ($secureToken.Length -eq 0) {
    Write-Error "Token cannot be empty."
    exit 1
}

$pointer = [Runtime.InteropServices.Marshal]::SecureStringToBSTR($secureToken)
try {
    $token = [Runtime.InteropServices.Marshal]::PtrToStringBSTR($pointer)

    Write-Host "Verifying token against GitHub API..." -ForegroundColor Cyan
    try {
        $user = Invoke-RestMethod `
            -Uri "https://api.github.com/user" `
            -Headers @{ 
                "Authorization" = "Bearer $token"
                "User-Agent"    = "MCP-Token-Setup"
                "Accept"        = "application/vnd.github+json"
            }
        Write-Host " Authenticated as: $($user.login) ($($user.name))" -ForegroundColor Green
    } catch {
        Write-Warning "Could not verify token online (Status: $($_.Exception.Message)). Proceeding with local configuration..."
    }

    # 1. Set Windows User Environment Variable
    [Environment]::SetEnvironmentVariable("GITHUB_PERSONAL_ACCESS_TOKEN", $token, "User")
    [Environment]::SetEnvironmentVariable("GITHUB_TOKEN", $token, "User")
    Write-Host " Saved GITHUB_PERSONAL_ACCESS_TOKEN and GITHUB_TOKEN to User Environment." -ForegroundColor Green

    # 2. Update Antigravity MCP Config (~/.gemini/config/mcp_config.json)
    $geminiConfig = "$HOME\.gemini\config\mcp_config.json"
    if (Test-Path $geminiConfig) {
        $content = Get-Content $geminiConfig -Raw
        $content = $content -replace '"GITHUB_PERSONAL_ACCESS_TOKEN":\s*"[^"]*"', "`"GITHUB_PERSONAL_ACCESS_TOKEN`": `"$token`""
        Set-Content -Path $geminiConfig -Value $content -Encoding UTF8
        Write-Host " Updated Antigravity MCP configuration: $geminiConfig" -ForegroundColor Green
    }

    # 3. Update OpenCode MCP Config (~/.config/opencode/opencode.jsonc)
    $opencodeConfig = "$HOME\.config\opencode\opencode.jsonc"
    if (Test-Path $opencodeConfig) {
        $content = Get-Content $opencodeConfig -Raw
        $content = $content -replace '"GITHUB_PERSONAL_ACCESS_TOKEN":\s*"[^"]*"', "`"GITHUB_PERSONAL_ACCESS_TOKEN`": `"$token`""
        Set-Content -Path $opencodeConfig -Value $content -Encoding UTF8
        Write-Host " Updated OpenCode MCP configuration: $opencodeConfig" -ForegroundColor Green
    }

    Write-Host "`n Setup completed successfully! Please restart your terminal/assistant to reload the new token." -ForegroundColor Cyan

} finally {
    [Runtime.InteropServices.Marshal]::ZeroFreeBSTR($pointer)
    $token = $null
}
