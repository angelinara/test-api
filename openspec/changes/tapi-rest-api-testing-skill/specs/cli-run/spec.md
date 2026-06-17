## ADDED Requirements

### Requirement: Run named request
`tapi run <name>` SHALL execute the request in `.test-api/requests/<name>.sh`.

#### Scenario: Named request found and executed
- **WHEN** user runs `tapi run login` and `login.sh` exists
- **THEN** the request is executed and output is displayed

#### Scenario: Named request not found
- **WHEN** user runs `tapi run login` and `login.sh` does not exist
- **THEN** an error message is printed listing available requests

### Requirement: Interactive picker when no name given
`tapi run` with no argument SHALL display an arrow-key selection list of all scripts in `.test-api/requests/`, showing name and description.

#### Scenario: Picker displayed
- **WHEN** user runs `tapi run` with no argument and requests exist
- **THEN** an interactive list is shown with each entry as `<name>  <description>`

#### Scenario: No requests available
- **WHEN** user runs `tapi run` with no argument and `.test-api/requests/` is empty
- **THEN** a message is printed suggesting `tapi new` to create a request

### Requirement: Request display block
Before executing, `tapi run` SHALL print a formatted REQUEST block showing method, URL, headers, and body parsed from the script.

#### Scenario: Request block printed
- **WHEN** request is about to execute
- **THEN** output shows `── REQUEST ──` header, method and URL on one line, each header on its own line, a blank line, then the body (if present)

### Requirement: Response display block
After executing, `tapi run` SHALL print a formatted RESPONSE block showing HTTP status code, elapsed time, and the response body piped through `jq`.

#### Scenario: Successful JSON response
- **WHEN** the server returns a 200 with a JSON body
- **THEN** output shows `── RESPONSE  200  <Xms> ──` header followed by jq-formatted body

#### Scenario: Non-JSON response body
- **WHEN** the server returns a body that is not valid JSON
- **THEN** body is printed as-is without jq formatting, with a note that it is not JSON

#### Scenario: Non-2xx status
- **WHEN** server returns a 4xx or 5xx status
- **THEN** status code is shown in the response header and body is displayed normally

### Requirement: jq prerequisite check
`tapi run` SHALL verify that `jq` is available in PATH before executing.

#### Scenario: jq not found
- **WHEN** `jq` is not in PATH
- **THEN** a clear error message is printed instructing the user to install jq, and the command exits without making a network request
