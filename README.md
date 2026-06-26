# tapi

A CLI and Claude Code skill for testing REST APIs. Requests are saved as bash scripts so they run standalone without `tapi` installed.

```bash
# install Go and jq
brew install go jq

# install tapi
go install github.com/angelinara/test-api/cmd/tapi

# set up tapi in your project (creates .test-api/requests/ and copies skills)
tapi init

# list saved requests
tapi list
```

## Usage

| Command | What it does |
|---------|-------------|
| `tapi init` | Set up `.test-api/requests/` and install Claude Code skills |
| `tapi list` | List saved requests with method, URL, and description |
| `/tapi-new` | Find routes in the codebase, create a new request file |
| `/tapi-list` | Pick a saved request and run it |

## Architecture

```mermaid
flowchart TD
    A[tapi init] --> B[Create .test-api/requests/]
    A --> C[Copy skills to .claude/skills/]

    D["/tapi-new"] --> E[Read source files for routes]
    E --> F[Run tapi list to find unsaved routes]
    F --> G[User picks a route]
    G --> H[Write .sh file to .test-api/requests/]

    I["/tapi-list"] --> J[Run tapi list]
    J --> K[User picks a request]
    K --> L[Run bash .test-api/requests/name.sh]
    L --> M[Display response]
```
