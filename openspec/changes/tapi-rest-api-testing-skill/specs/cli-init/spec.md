## ADDED Requirements

### Requirement: Initialize project directory
`tapi init` SHALL create the `.test-api/requests/` directory structure in the current working directory.

#### Scenario: Fresh project
- **WHEN** user runs `tapi init` in a directory with no `.test-api/` folder
- **THEN** `.test-api/requests/` is created and a success message is printed

#### Scenario: Already initialized
- **WHEN** user runs `tapi init` in a directory where `.test-api/requests/` already exists
- **THEN** command exits cleanly with a message indicating it is already initialized, making no changes
