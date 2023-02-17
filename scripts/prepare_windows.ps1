Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
$PSDefaultParameterValues['*:ErrorAction']='Stop'
function ThrowOnNativeFailure {
    if (-not $?)
    {
        throw 'Native Failure'
    }
}

Invoke-WebRequest -useb 'https://raw.githubusercontent.com/scoopinstaller/install/master/install.ps1' -outfile 'install.ps1'
.\install.ps1 -RunAsAdmin

Join-Path (Resolve-Path ~).Path "scoop\shims" >> $Env:GITHUB_PATH

scoop install jq
scoop install zip
scoop install curl
scoop install 7zip
scoop install go
