## ADDED Requirements

### Requirement: Read codebase and present endpoint list
The `/test-api` skill SHALL read the project source files directly to identify route definitions, cross-reference with `.test-api/requests/`, and present a merged list for the user to choose from.

#### Scenario: Endpoints found
- **WHEN** user invokes `/test-api` in a project with route definitions in source files
- **THEN** skill displays a list showing saved status, method, path, and for saved entries the request name and description

#### Scenario: No endpoints found
- **WHEN** skill finds no route definitions in the source files
- **THEN** skill informs the user no routes were detected and asks the user to describe the endpoint they want to test

### Requirement: Run saved request
When the user selects a saved endpoint, the skill SHALL execute it via bash.

#### Scenario: User selects saved endpoint
- **WHEN** user picks an endpoint marked as saved
- **THEN** skill runs `bash .test-api/requests/<name>.sh` and displays the output

### Requirement: Create then run unsaved request
When the user selects an unsaved endpoint, the skill SHALL gather request details conversationally, write the `.sh` file directly, then run it.

#### Scenario: User selects unsaved endpoint
- **WHEN** user picks an endpoint not yet saved
- **THEN** skill asks for any missing details (headers, body), writes `.test-api/requests/<name>.sh` with shebang on line 1, description comment on line 2, canonical curl flags, and `| jq .` at the end, sets permissions to 0755, then runs it via bash
