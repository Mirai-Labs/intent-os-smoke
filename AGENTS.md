# do not edit — managed by intent-os (intent: edd6c7ad-7b59-46ee-875f-b42dccc61b36, version: v1)
# AGENTS.md — Add a generic Set type to the smoke library

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Add a new file set.go to package smoke that declares the generic type alias `type Set[T comparable] = map[T]struct{}` (a generic type alias) plus the generic functions NewSet, SetAdd, SetHas and SetLen over it, and set_test.go with table-driven tests. Do not change calc.go or calc_test.go.

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

- set.go declares `type Set[T comparable] = map[T]struct{}` as a generic type alias (not a defined type) and NewSet[T comparable](items ...T) Set[T]
- set.go exports SetAdd, SetHas and SetLen as generic functions over Set[T]
- set_test.go has table-driven tests for NewSet, SetAdd, SetHas and SetLen with int and string element types, and go test ./... passes
