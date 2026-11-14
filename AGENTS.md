# do not edit — managed by intent-os (intent: 12a64873-84f3-4161-b24b-eacd2ba87d41, version: v1)
# AGENTS.md — Add Clamp to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Add Clamp(v, lo, hi int) int to calc.go returning v limited to the range lo..hi, with table-driven tests in calc_test.go. Do not change existing functions.

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

- calc.go exports Clamp(v, lo, hi int) int that returns lo when v<lo, hi when v>hi and v otherwise
- Clamp(5,0,3) returns 3, Clamp(-1,0,3) returns 0 and Clamp(2,0,3) returns 2
- calc_test.go has table-driven tests for Clamp and go test ./... passes
