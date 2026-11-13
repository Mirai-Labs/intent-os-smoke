# do not edit — managed by intent-os (intent: a90051f4-607b-446f-8f4e-8996352c435e, version: v1)
# AGENTS.md — Add Max and Min functions to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Extend package smoke in calc.go with Max(a, b int) int and Min(a, b int) int, with table-driven tests in calc_test.go. Do not change Add.

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

- calc.go exports Max(a, b int) int that returns the larger of a and b
- calc.go exports Min(a, b int) int that returns the smaller of a and b
- calc_test.go has table-driven tests for Max and Min and go test ./... passes
