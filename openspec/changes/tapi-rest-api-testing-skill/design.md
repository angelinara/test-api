## Context

No existing tooling in this repo. We are building `tapi` from scratch as a Go CLI + Claude Code skill. The tool is designed to be installed once and used across many projects. Request files are bash scripts so they remain portable and runnable without `tapi` installed.

## Goals / Non-Goals

**Goals:**
- Single binary (`tapi`) covering init, new, run, scan
- Requests as self-contained bash scripts with embedded description comment
- Formatted request + response output, body piped through `jq`
- Arrow-key picker when `tapi run` is called with no argument
- Claude Code skill that reads codebase routes and delegates to CLI
- Framework route detection: Gin, Express, FastAPI, Rails, Spring

**Non-Goals:**
- Response assertion / test suite runner (no pass/fail, no CI mode)
- Auth credential management or env file support
- Request chaining or variable extraction between requests
- GUI or web interface
- Collection import/export (Postman, Insomnia)

## Decisions

**Bash scripts over YAML/JSON**
Request files are `.sh` scripts, not structured data. Rationale: they are runnable standalone without `tapi`, diffable in git, and trivially copy-pasteable. The trade-off is that `tapi run` must parse curl flags at runtime — acceptable given the fixed shape `tapi new` generates.

**Description as bash comment, not metadata block**
`# Description here` on line 2 (after shebang) keeps the file valid bash with no special syntax. `tapi new` enforces ≤50 chars. `tapi` reads it by scanning line 2. Alternative (YAML frontmatter) would break standalone execution.

**Sentinel-based response splitting**
`tapi run` appends `-w "\n__TAPI__%{http_code}|%{time_total}"` when executing curl and splits stdout on `__TAPI__`. This avoids a second HTTP call for metadata. Risk: if response body contains `__TAPI__` literally, splitting breaks — accepted as negligible in practice.

**Go packages: cobra + huh**
`cobra` for subcommand routing (standard, well-understood). `huh` (charmbracelet) for interactive prompts and arrow-key picker — single dependency covering both `tapi new` (form) and `tapi run` (select). Alternative: `bubbletea` directly is more flexible but more boilerplate.

**`tapi scan` as JSON to stdout**
Route detection output is JSON so the skill can parse it with `jq` or read it directly. No custom format to maintain. The skill calls `tapi scan` and receives a stable contract.

**`internal/parser` parses curl flags, not the script**
`tapi run` parses the curl invocation from the `.sh` file to reconstruct method/url/headers/body for the REQUEST display block. It then re-invokes curl (not the script) so it can append `-w`. Alternative (executing script then capturing) would require intercepting stdin/stdout of a subprocess that calls curl — harder to control.

## Risks / Trade-offs

`jq` not installed → `tapi run` fails silently or with confusing error. Mitigation: check for `jq` in PATH at startup and print a clear prerequisite message.

Curl flag parsing is fragile if a user hand-edits the `.sh` file into non-standard forms. Mitigation: `tapi new` generates a canonical format; document that hand-edits should follow the same shape.

Route scanner regex is heuristic — will miss dynamic route registration and produce false positives on comments. Mitigation: output includes `file` and `line` so the user can verify. No auto-action is taken on scan results alone.

## Migration Plan

No existing state to migrate. Installation is `go install`. Per-project setup is `tapi init`. No rollback complexity.

## Open Questions

- Should `tapi run` fall back to executing the `.sh` script directly if curl flag parsing fails? (Currently: error out with a helpful message.)
- Multi-line body input in `tapi new`: `huh` textarea or line-by-line prompts? (Leaning toward textarea.)
