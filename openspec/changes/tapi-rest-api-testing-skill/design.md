## Context

No existing tooling in this repo. We are building `tapi` from scratch as a Go CLI + Claude Code skill. The tool is designed to be installed once and used across many projects. Request files are bash scripts so they remain portable and runnable without `tapi` installed.

## Goals / Non-Goals

**Goals:**
- Single binary (`tapi`) covering `init` and `list`
- Requests as self-contained bash scripts with embedded description comment and jq piped inline
- Claude Code skill that reads codebase routes, creates request files, and runs them
- Framework route detection: Gin, Express, FastAPI, Rails, Spring

**Non-Goals:**
- Response assertion / test suite runner (no pass/fail, no CI mode)
- Auth credential management or env file support
- Request chaining or variable extraction between requests
- GUI or web interface
- Collection import/export (Postman, Insomnia)

## Decisions

**Bash scripts over YAML/JSON**
Request files are `.sh` scripts, not structured data. Rationale: they are runnable standalone without `tapi`, diffable in git, and trivially copy-pasteable. The trade-off is that `tapi run` must parse curl flags at runtime — acceptable given the fixed shape the skill generates.

**Description as bash comment, not metadata block**
`# Description here` on line 2 (after shebang) keeps the file valid bash with no special syntax. The skill enforces ≤50 chars when creating files. `tapi` reads it by scanning line 2. Alternative (YAML frontmatter) would break standalone execution.

**Skill handles create and run — CLI handles init and list**
`tapi new` and `tapi run` are not CLI commands. The skill writes `.sh` files directly and executes them via `bash`. The CLI's responsibility is limited to `init` (set up the directory) and `list` (enumerate saved requests with names and descriptions). This keeps the binary minimal and puts conversational logic where it belongs — in the skill.

**jq piped inline in the `.sh` file**
Request scripts pipe curl output through `jq .` directly: `curl -s ... | jq .`. The skill writes this format. No sentinel or response-splitting logic needed in the CLI.

**No external dependencies**
With only 2 subcommands (`init`, `list`), a `switch os.Args[1]` in `main.go` is sufficient. Zero external dependencies.

**No scanner — skill reads source directly**
Route detection is done by the skill (Claude) reading source files directly — no `tapi scan` command, no `internal/scanner` package. Claude is better at this than regex heuristics.

**`internal/parser` parses curl flags and description**
`tapi run` prints the raw `.sh` file contents as the request display. The parser's job is to extract the description from line 2 for the picker, and to extract method/url/headers/body so the runner can re-invoke curl with the sentinel `-w` flag appended.

## Risks / Trade-offs

`jq` not installed → `tapi run` fails silently or with confusing error. Mitigation: check for `jq` in PATH at startup and print a clear prerequisite message.

Curl flag parsing is fragile if a user hand-edits the `.sh` file into non-standard forms. Mitigation: the skill generates a canonical format; document that hand-edits should follow the same shape.

Route scanner regex is heuristic — will miss dynamic route registration and produce false positives on comments. Mitigation: output includes `file` and `line` so the user can verify. No auto-action is taken on scan results alone.

## Migration Plan

No existing state to migrate. Installation is `go install`. Per-project setup is `tapi init`. No rollback complexity.

## Open Questions

- Should `tapi run` fall back to executing the `.sh` script directly if curl flag parsing fails? (Currently: error out with a helpful message.)
