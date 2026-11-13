# do not edit — managed by intent-os (intent: 29a2941a-652e-4224-8b42-e9d8a99cc248, version: v1)
# AGENTS.md — Add Sub and Mul functions to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Extend package smoke in calc.go with Sub(a, b int) int and Mul(a, b int) int, with table-driven tests in calc_test.go. Do not change Add.

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

- calc.go exports Sub(a, b int) int that returns a minus b
- calc.go exports Mul(a, b int) int that returns a times b
- calc_test.go has table-driven tests for Sub and Mul and go test ./... passes
