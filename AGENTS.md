# do not edit — managed by intent-os (intent: cc760661-70af-42e9-b582-7131f60dbfbf, version: v1)
# AGENTS.md — Add Double and Half functions to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Extend package smoke in calc.go with Double(a int) int and Half(a int) int, with table-driven tests in calc_test.go. Do not change Add.

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

- calc.go exports Double(a int) int that returns a times 2
- calc.go exports Half(a int) int that returns a divided by 2, rounding toward zero
- calc_test.go has table-driven tests for Double and Half and go test ./... passes
