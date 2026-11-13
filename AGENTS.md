# do not edit — managed by intent-os (intent: 0eae61fb-0d99-4634-a0d9-dc92326a7187, version: v1)
# AGENTS.md — Add Negate and a CI workflow to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Add Negate(a int) int to calc.go with table-driven tests in calc_test.go. Do not change Add. Also add a GitHub Actions workflow at .github/workflows/test.yml that runs go test ./... on every push.

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

- calc.go exports Negate(a int) int that returns minus a
- Negate(0) returns 0 and Negate(Negate(x)) returns x
- calc_test.go has table-driven tests for Negate and go test ./... passes
