# do not edit — managed by intent-os (intent: 4093edd6-8246-4fe2-9d25-8ca8dda5ee85, version: v1)
# AGENTS.md — Add Clamp and Sign functions to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Extend package smoke in calc.go with Clamp(v, lo, hi int) int and Sign(a int) int, with table-driven tests in calc_test.go. Do not change Add.

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

- calc.go exports Clamp(v, lo, hi int) int that limits v to the range lo..hi
- calc.go exports Sign(a int) int that returns -1, 0 or 1 according to the sign of a
- calc_test.go has table-driven tests for Clamp and Sign and go test ./... passes
