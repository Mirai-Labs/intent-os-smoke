# do not edit — managed by intent-os (intent: 8effa0c8-7ee1-4d3f-a8c7-846298d5ef10, version: v1)
# AGENTS.md — Add Abs to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Add Abs(a int) int to calc.go returning the absolute value, with table-driven tests in calc_test.go. Do not change existing functions.

## Artifacts Under Management

- `.controlplane/manifest.yaml`
- `.controlplane/context-summary.md`
- `.controlplane/constraints.md`
- `.controlplane/acceptance.md`
- `.controlplane/commands.md`
- `.controlplane/memory.md`
- `AGENTS.md`
- `CLAUDE.md`
- `.cursor/rules/intent-os.mdc`
- `.claude/settings.json`
- `.controlplane/hooks/pretooluse.sh`
- `.controlplane/hooks/posttooluse.sh`
- `.claude/mcp_servers.json`
- `.controlplane/PRD.md`
- `.controlplane/TASK.md`
- `.controlplane/MEMORY.snapshot.md`
- `.controlplane/context-map.json`
- `.controlplane/guardrails.json`
- `.github/copilot-instructions.md`

## Required Checks

- lint
- test
- policy:secret-leak

## Constraints

Allowed tools: `read`, `write`, `grep`, `run`

Denied tools: `network_egress`

## Acceptance Criteria

- calc.go exports Abs(a int) int that returns the absolute value of a
- Abs(-3) returns 3 and Abs(3) returns 3 and Abs(0) returns 0
- calc_test.go has table-driven tests for Abs and go test ./... passes
