## 1. Project Scaffold

- [x] 1.1 Initialize Go module (`go mod init`) — no external dependencies required
- [x] 1.2 Create directory structure: `cmd/tapi/`, `internal/parser/`, `internal/runner/`, `internal/formatter/`, `internal/picker/`, `internal/builder/`, `internal/scanner/`
- [x] 1.3 Implement subcommand routing in `cmd/tapi/main.go` via `switch os.Args[1]` for `init`, `new`, `run`, `scan`

## 2. tapi init

- [x] 2.1 Implement `init` command: create `.test-api/requests/` in cwd
- [x] 2.2 Handle already-initialized case: check existence and print appropriate message

## 3. internal/parser

- [ ] 3.1 Implement curl flag parser: extract `-X METHOD`, URL, `-H` headers, `-d` body from a `.sh` file
- [ ] 3.2 Read description from line 2 (first comment after shebang)
- [ ] 3.3 Write unit tests covering GET (no body), POST with headers and body, multi-line body

## 4. internal/scanner

- [ ] 4.1 Implement recursive file walker, skipping `vendor/`, `node_modules/`, `.git/`
- [ ] 4.2 Add regex patterns for Gin (`r.GET`, `r.POST`, `r.PUT`, `r.PATCH`, `r.DELETE`)
- [ ] 4.3 Add regex patterns for Express (`app.get`, `app.post`, `router.get`, `router.post`, etc.)
- [ ] 4.4 Add regex patterns for FastAPI (`@app.get`, `@router.post`, etc.)
- [ ] 4.5 Add regex patterns for Rails (`get '/path'`, `post '/path'`, etc.)
- [ ] 4.6 Add regex patterns for Spring (`@GetMapping`, `@PostMapping`, `@PutMapping`, `@DeleteMapping`, `@PatchMapping`)
- [ ] 4.7 Cross-reference `.test-api/requests/` to populate `saved`, `name`, `description` fields
- [ ] 4.8 Implement `scan` command: run scanner, marshal to JSON, write to stdout

## 5. internal/builder (tapi new)

- [x] 5.1 Parse flags from args: `--name`, `--description`, `--method`, `--url`, `--header` (repeatable), `--body`
- [x] 5.2 Validate required fields: name, description (≤50 chars), method, url; exit with error if missing
- [ ] 5.3 Implement script writer: generate `.sh` with shebang, description comment, canonical curl command, set permissions to 0755
- [ ] 5.4 Handle name collision: exit with error if file already exists (skill handles overwrite logic)
- [ ] 5.5 Wire builder into `new` command

## 6. internal/runner + internal/formatter (tapi run)

- [ ] 6.1 Implement jq prerequisite check: verify `jq` in PATH, exit with clear message if missing
- [ ] 6.2 Implement request execution: parse script via `internal/parser`, re-invoke curl appending `-w "\n__TAPI__%{http_code}|%{time_total}"`
- [ ] 6.3 Implement stdout splitter: split on `__TAPI__` sentinel, separate body from status+timing
- [ ] 6.4 Implement body formatter: pipe through `jq .`; fall back to raw print with note if jq exits non-zero
- [ ] 6.5 Implement REQUEST display block: `── REQUEST ──` header, method + URL, headers, blank line, body
- [ ] 6.6 Implement RESPONSE display block: `── RESPONSE  <status>  <Xms> ──` header, formatted body
- [ ] 6.7 Wire runner + formatter into `run` command with name argument

## 7. internal/picker (tapi run — no argument)

- [ ] 7.1 Implement request lister: read `.test-api/requests/*.sh`, extract name + description via parser
- [ ] 7.2 Print numbered list of requests showing `<number>) <name>  <description>`, read selection from stdin
- [ ] 7.3 Handle empty requests directory: print message suggesting `/test-api` skill
- [ ] 7.4 Wire picker into `run` command when no name argument is provided

## 8. Claude Code Skill

- [ ] 8.1 Create `.claude/commands/test-api.md` skill file
- [ ] 8.2 Skill checks for `tapi` in PATH; prints install instructions if missing
- [ ] 8.3 Skill runs `tapi scan`, parses JSON output
- [ ] 8.4 Skill presents merged endpoint list (saved/unsaved, method, path, name, description)
- [ ] 8.5 Skill runs `tapi run <name>` for saved endpoint selections
- [ ] 8.6 Skill invokes `tapi new` then `tapi run <name>` for unsaved endpoint selections

## 9. README and Polish

- [ ] 9.1 Write `README.md`: prerequisites (Go, jq), install (`go install`), per-project setup (`tapi init`), usage (`tapi new`, `tapi run`, `/test-api` skill)
- [ ] 9.2 Add `Makefile` target: `make install` runs `go install ./cmd/tapi/`
- [ ] 9.3 Smoke test: run `tapi init`, `tapi new`, `tapi run`, `tapi scan` against a sample project
