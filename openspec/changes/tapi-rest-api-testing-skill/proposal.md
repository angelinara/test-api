## Why

Testing REST API endpoints during development requires switching between tools, managing state across files, and lacks a conversational interface. `tapi` brings endpoint testing into the terminal and Claude Code workflow — requests live as plain bash scripts in the repo, and a skill discovers endpoints automatically.

## What Changes

- New Go CLI binary (`tapi`) with `init` and `run` commands
- Requests persisted as self-contained bash scripts under `.test-api/requests/`
- Claude Code skill (`/test-api`) that reads the codebase for route definitions, creates request files directly, and delegates execution to the CLI
- Numbered picker and formatted request/response output with jq

## Capabilities

### New Capabilities

- `cli-init`: Initialize `.test-api/requests/` directory in any project
- `cli-run`: Execute a saved request, display formatted request + response piped through jq
- `test-api-skill`: Claude Code skill that reads source files for routes, creates request files directly, and delegates execution to `tapi run`

### Modified Capabilities

## Impact

- New Go module in this repo (`cmd/tapi/`, `internal/`)
- Installs `tapi` binary to `$GOPATH/bin` or via `go install`
- Requires `jq` on the host system at runtime
- Per-project: `.test-api/requests/` added to target repos; `.gitignore` not modified (request scripts should be committed)
- `.claude/commands/test-api.md` ships with this repo for users to copy into their projects
