# Estoty Nakama Config Service

Dockerized Nakama + PostgreSQL project with three Go runtime RPC methods:

- `update_metadata`: merges caller supplied JSON into the authenticated user's account metadata.
- `config`: returns the free-form JSON game configuration from `config/game_config.json`.
- `private_ping`: server-to-server RPC that succeeds only when called with the Nakama runtime HTTP key.

## Requirements

- Docker Desktop
- Docker Compose v2

## Start

```powershell
docker compose up --build
```

Nakama HTTP API runs on `http://127.0.0.1:7350`.
Nakama console runs on `http://127.0.0.1:7351`.

Development keys used by this compose file:

- Socket server key: `defaultkey`
- Server-to-server HTTP key: `s2s-dev-key`

## Call the RPCs

Create a test user session:

```powershell
$auth = "Basic " + [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes("defaultkey:"))
$sessionBody = @{ id = "player-1" } | ConvertTo-Json -Compress
$session = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:7350/v2/account/authenticate/custom?create=true" -Headers @{ Authorization = $auth } -ContentType "application/json" -Body $sessionBody
$token = $session.token
```

Update user metadata:

```powershell
$payload = '{"coins":100,"favorite_color":"blue"}' | ConvertTo-Json -Compress
Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:7350/v2/rpc/update_metadata" -Headers @{ Authorization = "Bearer $token" } -ContentType "application/json" -Body $payload
```

Read game configuration:

```powershell
$payload = '{}' | ConvertTo-Json -Compress
Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:7350/v2/rpc/config" -Headers @{ Authorization = "Bearer $token" } -ContentType "application/json" -Body $payload
```

Call the private server-to-server RPC:

```powershell
$payload = '{}' | ConvertTo-Json -Compress
Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:7350/v2/rpc/private_ping?http_key=s2s-dev-key" -ContentType "application/json" -Body $payload
```

Calling `private_ping` with a user session is rejected:

```powershell
$payload = '{}' | ConvertTo-Json -Compress
Invoke-WebRequest -UseBasicParsing -Method Post -Uri "http://127.0.0.1:7350/v2/rpc/private_ping" -Headers @{ Authorization = "Bearer $token" } -ContentType "application/json" -Body $payload
```

## Notes

This project pins `heroiclabs/nakama:3.22.0` and `github.com/heroiclabs/nakama-common v1.32.0` so the Go plugin ABI matches the Nakama server version.
