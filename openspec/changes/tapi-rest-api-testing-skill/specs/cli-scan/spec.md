## ADDED Requirements

### Requirement: Walk source files for route definitions
`tapi scan` SHALL recursively walk the current directory and detect HTTP route definitions using regex patterns for supported frameworks.

#### Scenario: Routes detected
- **WHEN** source files contain route definitions in a supported framework
- **THEN** each route is included in the output with method, path, file, and line number

#### Scenario: No routes found
- **WHEN** no route definitions are detected in the project
- **THEN** an empty JSON array is output

### Requirement: Supported frameworks
`tapi scan` SHALL detect route patterns for: Gin, Express, FastAPI, Rails, and Spring.

#### Scenario: Gin routes detected
- **WHEN** source contains `r.GET("/path"`, `r.POST("/path"`, etc.
- **THEN** method and path are extracted correctly

#### Scenario: Express routes detected
- **WHEN** source contains `app.get('/path'`, `router.post('/path'`, etc.
- **THEN** method and path are extracted correctly

#### Scenario: FastAPI routes detected
- **WHEN** source contains `@app.get("/path"`, `@router.post("/path"`, etc.
- **THEN** method and path are extracted correctly

#### Scenario: Rails routes detected
- **WHEN** source contains `get '/path'`, `post '/path'`, etc. in a routes file
- **THEN** method and path are extracted correctly

#### Scenario: Spring routes detected
- **WHEN** source contains `@GetMapping("/path"`, `@PostMapping("/path"`, etc.
- **THEN** method and path are extracted correctly

### Requirement: Cross-reference saved requests
`tapi scan` SHALL check `.test-api/requests/` and mark each route as saved or unsaved, including the request name and description when saved.

#### Scenario: Saved route
- **WHEN** a detected route has a matching `.sh` file in `.test-api/requests/`
- **THEN** output includes `"saved": true`, `"name": "<name>"`, and `"description": "<description>"`

#### Scenario: Unsaved route
- **WHEN** a detected route has no matching request file
- **THEN** output includes `"saved": false` and omits name and description

### Requirement: JSON output to stdout
`tapi scan` SHALL write a JSON array to stdout, one object per detected route.

#### Scenario: Output schema
- **WHEN** routes are detected
- **THEN** each object contains: `method` (string), `path` (string), `file` (string, relative path), `line` (int), `saved` (bool), and optionally `name` and `description`
