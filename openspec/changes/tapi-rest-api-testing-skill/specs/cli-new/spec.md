## ADDED Requirements

### Requirement: Interactive request builder
`tapi new [name]` SHALL interactively prompt the user for all fields required to create a request and save it as a bash script.

#### Scenario: Create with name argument
- **WHEN** user runs `tapi new login`
- **THEN** interactive prompts begin with name pre-filled as `login`

#### Scenario: Create without name argument
- **WHEN** user runs `tapi new` with no argument
- **THEN** first prompt asks for the request name

### Requirement: Required description
The description prompt SHALL be required and enforce a maximum of 50 characters.

#### Scenario: Empty description rejected
- **WHEN** user submits an empty description
- **THEN** prompt displays a validation error and re-prompts

#### Scenario: Description over 50 chars rejected
- **WHEN** user enters a description longer than 50 characters
- **THEN** prompt displays a validation error and re-prompts

### Requirement: Method selection
The method prompt SHALL present GET, POST, PUT, PATCH, DELETE as a select list.

#### Scenario: Method selected
- **WHEN** user selects POST from the method list
- **THEN** POST is recorded and the URL prompt follows

### Requirement: Header collection
The header prompt SHALL allow zero or more headers to be added as `key: value` pairs, stopping when the user submits an empty entry.

#### Scenario: No headers
- **WHEN** user submits an empty entry on the first header prompt
- **THEN** no headers are recorded and the prompt moves on

#### Scenario: Multiple headers
- **WHEN** user enters `Content-Type: application/json` then `Authorization: Bearer token` then empty
- **THEN** both headers are recorded

### Requirement: Body input
The body prompt SHALL appear only for POST, PUT, and PATCH methods and accept multi-line text input.

#### Scenario: Body skipped for GET
- **WHEN** user selects GET as the method
- **THEN** body prompt is not shown

#### Scenario: Body accepted for POST
- **WHEN** user selects POST and enters a JSON body
- **THEN** body text is recorded verbatim

### Requirement: Script output format
`tapi new` SHALL write the request to `.test-api/requests/<name>.sh` in a canonical format parseable by `tapi run`.

#### Scenario: Script is written
- **WHEN** all prompts are completed
- **THEN** `.test-api/requests/<name>.sh` is created with shebang on line 1, description comment on line 2, and a curl command with `-s` flag, `-X METHOD`, all `-H` headers, optional `-d` body, and URL as final argument

#### Scenario: Script is executable
- **WHEN** script is written
- **THEN** file permissions are set to executable (0755)

#### Scenario: Name collision
- **WHEN** a script with the given name already exists
- **THEN** user is prompted to confirm overwrite before proceeding
