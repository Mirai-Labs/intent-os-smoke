# do not edit — managed by intent-os (intent: 93ad5f83-66cd-4b2c-8b56-122a45890c58, version: v1)
# AGENTS.md — Add Median to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Extend package smoke in calc.go with Median(xs []int) float64, with table-driven tests in calc_test.go. Do not change Add.

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

- calc.go exports Median(xs []int) float64: the middle value of the sorted values, the mean of the two middle values for an even count, and 0 for an empty slice
- Median never modifies or reorders the slice passed by the caller, and calc_test.go proves it with a test that compares the input before and after
- calc_test.go has table-driven tests for odd, even, empty and unsorted inputs, and go test ./... passes
