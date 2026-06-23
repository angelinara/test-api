## 1. Project Scaffold

- [x] 1.1 Initialize Go module (`go mod init`) — no external dependencies required
- [x] 1.2 Create directory structure: `cmd/tapi/`, `internal/parser/`
- [x] 1.3 Implement subcommand routing in `cmd/tapi/main.go` via `switch os.Args[1]` for `init`, `list`

## 2. tapi init

- [x] 2.1 Implement `init` command: create `.test-api/requests/` in cwd
- [x] 2.2 Handle already-initialized case: check existence and print appropriate message

## 3. internal/parser

- [x] 3.1 Implement curl flag parser: extract `-X METHOD`, URL, `-H` headers, `-d` body from a `.sh` file
- [x] 3.2 Read description from line 2 (first comment after shebang)
- [x] 3.3 Write unit tests covering GET (no body), POST with headers and body, multi-line body

## 4. tapi list

- [ ] 4.1 Read `.test-api/requests/*.sh`, extract name (filename without extension) and description (line 2) via parser
- [ ] 4.2 Print each request as `<name>  <description>`
- [ ] 4.3 Handle empty requests directory: print message suggesting `/test-api` skill
- [ ] 4.4 Wire list command into `main.go`

## 5. Claude Code Skill

- [ ] 5.1 Create `.claude/commands/test-api.md` skill file
- [ ] 5.2 Skill reads codebase source files directly to identify routes (method + path)
- [ ] 5.3 Skill reads `.test-api/requests/` to find already saved requests
- [ ] 5.4 Skill presents merged endpoint list (saved/unsaved, method, path, name, description)
- [ ] 5.5 Skill runs `bash .test-api/requests/<name>.sh` for saved endpoint selections and displays output
- [ ] 5.6 Skill writes `.sh` file directly to `.test-api/requests/<name>.sh` then runs it for unsaved endpoints

## 6. README and Polish

- [ ] 6.1 Write `README.md`: prerequisites (Go, jq), install (`go install`), per-project setup (`tapi init`), usage (`tapi list`, `/test-api` skill)
- [ ] 6.2 Smoke test: run `tapi init`, create a request via skill, `tapi list`, run via skill
