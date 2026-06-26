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

Both Claude skills shell out to the `tapi` binary — `tapi list` is the shared
boundary between the conversational layer and the CLI.

```mermaid
flowchart TD
    subgraph skills ["Claude Code skills (.claude/skills/)"]
        direction TB
        N["/tapi-new"] --> N1[Read source files for routes]
        N1 --> N2[User picks an unsaved route]
        N2 --> N3[Write .sh to .test-api/requests/]

        L["/tapi-list"] --> L1[User picks a saved request]
        L1 --> L2["bash .test-api/requests/&lt;name&gt;.sh"]
        L2 --> L3[Display response]
    end

    subgraph cli ["CLI commands (tapi binary)"]
        direction TB
        I["tapi init"] --> I1[Create .test-api/requests/]
        I --> I2[Copy skills to .claude/skills/]

        LS["tapi list"] --> LS1[Parse .sh files: name, method, URL, desc]
    end

    %% what calls what
    N1 -. "calls (exclude saved)" .-> LS
    L -. calls .-> LS
    I2 -. installs .-> skills
```
