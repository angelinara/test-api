---
name: tapi-new
description: Create a new saved API request for this project. Use when the user invokes /tapi-new, wants to add a new request, wants to test an endpoint they haven't saved yet, or asks to create a curl request for a route in the codebase. Always use this skill when the user wants to set up a new API request file, even if they don't use those exact words.
---

## What this skill does

Finds routes in the codebase that don't have saved requests yet, lets the user pick one, and writes a ready-to-run `.sh` file for it.

## Steps

### 1. Find routes in the codebase

Read the source files directly to identify API routes — look for method + path pairs. Support these frameworks:

- **Gin (Go)**: `r.GET(...)`, `r.POST(...)`, etc.
- **Express (Node)**: `app.get(...)`, `router.post(...)`, etc.
- **FastAPI (Python)**: `@app.get(...)`, `@router.post(...)`, etc.
- **Rails (Ruby)**: `get '...'`, `resources :...` in routes.rb
- **Spring (Java)**: `@GetMapping`, `@PostMapping`, etc.

List every route you find as `METHOD /path`.

### 2. Check what's already saved

Run:

```bash
tapi list
```

The output includes name, method, URL, and description for each saved request. Use the method and path to match against the routes you found — exclude any that are already covered.

If `tapi list` says no requests found, that's fine — all routes are unsaved.

### 3. Present unsaved routes

Show the user the routes that don't have saved requests yet and ask which one they'd like to create.

If all routes are already saved, tell the user:

> All routes already have saved requests. Use `/tapi-list` to run one.

### 4. Gather request details

Ask the user for anything you can't infer from the route:

- **Description** — one short line, 50 chars max (e.g. "Get all users", "Create a post")
- **Request body** — if it's a POST/PUT/PATCH, ask what fields to include
- **Headers** — ask if any are needed (e.g. auth token, content type)
- **Base URL** — if not obvious from the codebase (default to `http://localhost:8080` if unclear)

Don't ask for things you can already infer — method and path come from the route.

### 5. Write the `.sh` file

Write to `.test-api/requests/<name>.sh` where `<name>` is a short kebab-case identifier derived from the route (e.g. `get-users`, `create-post`).

Use this exact format:

```bash
#!/bin/bash
# <description>
curl -s \
  -X <METHOD> \
  -H "Content-Type: application/json" \
  -d '<body>' \
  "<base-url><path>" | jq .
```

Omit `-H` and `-d` lines if there are no headers or body. For GET requests, there's typically no body.

Make the file executable:

```bash
chmod +x .test-api/requests/<name>.sh
```

### 6. Confirm

Tell the user the file was saved and how to run it:

> Saved to `.test-api/requests/<name>.sh`. Run it with `/tapi-list` or `bash .test-api/requests/<name>.sh`.
