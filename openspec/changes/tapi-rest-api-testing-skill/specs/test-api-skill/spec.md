## ADDED Requirements

### Requirement: Read codebase and present endpoint list
The `/test-api` skill SHALL read the project source files directly to identify route definitions, cross-reference with `.test-api/requests/`, and present a merged list for the user to choose from.

#### Scenario: Endpoints found
- **WHEN** user invokes `/test-api` in a project with route definitions in source files
- **THEN** skill displays a list showing saved status, method, path, and for saved entries the request name and description

#### Scenario: No endpoints found
- **WHEN** skill finds no route definitions in the source files
- **THEN** skill informs the user no routes were detected and suggests running `tapi new` directly

### Requirement: Run saved request
When the user selects a saved endpoint, the skill SHALL execute `tapi run <name>`.

#### Scenario: User selects saved endpoint
- **WHEN** user picks an endpoint marked as saved
- **THEN** skill runs `tapi run <name>` and displays the output

### Requirement: Create then run unsaved request
When the user selects an unsaved endpoint, the skill SHALL invoke `tapi new` to create the request and then run it.

#### Scenario: User selects unsaved endpoint
- **WHEN** user picks an endpoint not yet saved
- **THEN** skill invokes `tapi new` with the endpoint name pre-suggested, waits for completion, then runs `tapi run <name>`

### Requirement: tapi CLI prerequisite check
The skill SHALL verify that the `tapi` binary is available before proceeding.

#### Scenario: tapi not installed
- **WHEN** `tapi` is not found in PATH
- **THEN** skill prints installation instructions (`go install`) and exits without running scan
