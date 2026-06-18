## REMOVED Requirements

### Requirement: Walk source files for route definitions
**Reason**: Route detection is handled by the skill (Claude) reading source files directly. Regex-based scanning in the CLI is redundant and less accurate than Claude reading the codebase.
**Migration**: Use the `/test-api` skill instead of `tapi scan`.
