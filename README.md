# ScriptStakingUnStaking


$env:GOOS = "darwin"
$env:GOARCH = "arm64"
go build -o batch-trans-macos-arm64 .\main.go

$env:GOOS = $null
$env:GOARCH = $null

