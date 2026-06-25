## 1. Project Scaffold

- [x] 1.1 Initialize Go module (`go mod init`) — no external dependencies required
- [x] 1.2 Create directory structure: `cmd/tapi/`, `internal/parser/`
- [x] 1.3 Implement subcommand routing in `cmd/tapi/main.go` via `switch os.Args[1]` for `init`, `list`

## 2. tapi init

- [x] 2.1 Implement `init` command: create `.test-api/requests/` in cwd
- [x] 2.2 Handle already-initialized case: check existence and print appropriate message
- [ ] 2.3 Embed skill files into binary using `//go:embed` (store under `internal/skills/`)
- [ ] 2.4 Copy `tapi-new` and `tapi-list` skill files into `.claude/skills/` in cwd on `tapi init`
- [ ] 2.5 Create `.claude/skills/` if it doesn't exist
- [ ] 2.6 Skip copy if skill files already exist, print message

## 3. internal/parser

- [x] 3.1 Implement curl flag parser: extract `-X METHOD`, URL, `-H` headers, `-d` body from a `.sh` file
- [x] 3.2 Read description from line 2 (first comment after shebang)
- [x] 3.3 Write unit tests covering GET (no body), POST with headers and body, multi-line body

## 4. tapi list

- [x] 4.1 Read `.test-api/requests/*.sh`, extract name (filename without extension) and description (line 2) via parser
- [x] 4.2 Print each request as `<name>  <description>`
- [x] 4.3 Handle empty requests directory: print message suggesting `/tapi-new` skill
- [x] 4.4 Wire list command into `main.go`
- [x] 4.5 Extend `ListItem` to include `Method` and `URL` via `ParseFile`
- [x] 4.6 Update `tapi list` output to include method and URL: `<name>  <method>  <url>  <description>`
- [x] 4.7 Update tests for new output format

## 5. Claude Code Skills

### /tapi-new
- [x] 5.1 Create `.claude/skills/tapi-new.md` skill file
- [x] 5.2 Skill reads codebase source files directly to identify routes (method + path)
- [x] 5.3 Skill runs `tapi list` to get saved requests (name, method, URL) and excludes matching routes
- [x] 5.4 Skill presents unsaved endpoints for selection
- [x] 5.5 Skill writes `.sh` file to `.test-api/requests/<name>.sh`

### /tapi-list
- [x] 5.6 Create `.claude/skills/tapi-list.md` skill file
- [x] 5.7 Skill runs `tapi list` to show saved requests
- [x] 5.8 Skill lets user pick a request and runs `bash .test-api/requests/<name>.sh`
- [x] 5.9 Skill displays output

## 6. README and Polish

- [x] 6.1 Write `README.md`: prerequisites (Go, jq), install (`go install github.com/angelinara/test-api/cmd/tapi@latest` — repo is public), per-project setup (`tapi init`), usage (`tapi list`, `/tapi-new`, `/tapi-list`)
- [x] 6.2 Smoke test: run `tapi init`, create a request via skill, `tapi list`, run via skill
