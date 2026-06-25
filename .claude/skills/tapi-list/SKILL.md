---
name: tapi-list
description: List and run saved API requests for this project. Use when the user invokes /tapi-list, wants to see saved requests, wants to run an existing request, or wants to test an API endpoint they've saved before.
---

## What this skill does

Shows the saved API requests for this project, lets the user pick one, runs it, and displays the response. If the user wants to run another after, loop back and show the list again.

## Steps

### 1. Check for saved requests

Run:
```bash
tapi list
```

If the output says no requests were found, tell the user:
> No saved requests yet. Run `/tapi-new` to create your first one.

Then stop.

### 2. Present the list

Show the output of `tapi list` to the user and ask which request they'd like to run.

### 3. Run the selected request

Run:
```bash
bash .test-api/requests/<name>.sh
```

### 4. Display the output

Show the response clearly. If the command fails (non-zero exit, connection refused, or no response), don't show the raw error — instead say something like:
> Could not reach the server. Make sure it's running locally and try again.

If `jq` is not installed, say:
> `jq` is required to format the response. Install it with `brew install jq` and try again.

### 5. Ask to run another

After showing the response, ask:
> Run another request?

If yes, go back to step 2. If no, stop.
