## 1. Project Scaffold

- [x] 1.1 Initialize Go module (`go mod init`) — no external dependencies required
- [x] 1.2 Create directory structure: `cmd/tapi/`, `internal/parser/`, `internal/runner/`, `internal/formatter/`, `internal/picker/`, `internal/builder/`
- [x] 1.3 Implement subcommand routing in `cmd/tapi/main.go` via `switch os.Args[1]` for `init`, `new`, `run`

## 2. tapi init

- [x] 2.1 Implement `init` command: create `.test-api/requests/` in cwd
- [x] 2.2 Handle already-initialized case: check existence and print appropriate message

## 3. internal/parser

- [x] 3.1 Implement curl flag parser: extract `-X METHOD`, URL, `-H` headers, `-d` body from a `.sh` file
- [x] 3.2 Read description from line 2 (first comment after shebang)
- [x] 3.3 Write unit tests covering GET (no body), POST with headers and body, multi-line body

## 4. internal/builder (tapi new)

- [x] 4.1 Parse flags from args: `--name`, `--description`, `--method`, `--url`, `--header` (repeatable), `--body`
- [x] 4.2 Validate required fields: name, description (≤50 chars), method, url; exit with error if missing
- [ ] 4.3 Implement script writer: generate `.sh` with shebang, description comment, canonical curl command, set permissions to 0755
- [ ] 4.4 Handle name collision: exit with error if file already exists (skill handles overwrite logic)
- [ ] 4.5 Wire builder into `new` command

## 5. internal/runner + internal/formatter (tapi run)

- [ ] 5.1 Implement jq prerequisite check: verify `jq` in PATH, exit with clear message if missing
- [ ] 5.2 Print raw `.sh` file contents as the request block before executing
- [ ] 5.3 Execute the script, appending `-w "\n__TAPI__%{http_code}|%{time_total}"` to the curl call
- [ ] 5.4 Split stdout on `__TAPI__` sentinel, separate body from status+timing
- [ ] 5.5 Pipe body through `jq .`; fall back to raw print with note if jq exits non-zero
- [ ] 5.6 Implement RESPONSE display block: `── RESPONSE  <status>  <Xms> ──` header, formatted body
- [ ] 5.7 Wire runner + formatter into `run` command with name argument

## 6. internal/picker (tapi run — no argument)

- [ ] 6.1 Read `.test-api/requests/*.sh`, extract name (filename) and description (line 2) via parser
- [ ] 6.2 Print numbered list showing `<number>) <name>  <description>`, read selection from stdin
- [ ] 6.3 Handle empty requests directory: print message suggesting `/test-api` skill
- [ ] 6.4 Wire picker into `run` command when no name argument is provided

## 7. Claude Code Skill

- [ ] 7.1 Create `.claude/commands/test-api.md` skill file
- [ ] 7.2 Skill checks for `tapi` in PATH; prints install instructions if missing
- [ ] 7.3 Skill reads codebase source files directly to identify routes (method + path)
- [ ] 7.4 Skill reads `.test-api/requests/` to find already saved requests
- [ ] 7.5 Skill presents merged endpoint list (saved/unsaved, method, path, name, description)
- [ ] 7.6 Skill runs `tapi run <name>` for saved endpoint selections
- [ ] 7.7 Skill invokes `tapi new ...` with all flags then `tapi run <name>` for unsaved endpoints

## 8. README and Polish

- [ ] 8.1 Write `README.md`: prerequisites (Go, jq), install (`go install`), per-project setup (`tapi init`), usage (`tapi new`, `tapi run`, `/test-api` skill)
- [ ] 8.2 Smoke test: run `tapi init`, `tapi new`, `tapi run` against a sample project
