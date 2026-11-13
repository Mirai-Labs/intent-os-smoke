# do not edit — managed by intent-os (intent: 5f6aa215-1fef-429e-89e7-ac811c84b332, version: v1)
# AGENTS.md — Add generic slice helpers (second attempt)

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Add a new file slices.go to package smoke with the generic helpers Map, Filter, Reduce, GroupBy and Zip, and slices_test.go with table-driven tests for each of them. Do not change calc.go or calc_test.go.

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

- slices.go exports Map[T, R any](xs []T, f func(T) R) []R, Filter[T any](xs []T, pred func(T) bool) []T and Reduce[T, A any](xs []T, init A, f func(A, T) A) A
- slices.go exports GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T that keeps the input order inside each group, and Zip[A, B any](as []A, bs []B) []struct{ First A; Second B } that stops at the shorter slice
- slices_test.go has table-driven tests for Map, Filter, Reduce, GroupBy and Zip using at least two different type parameters each, and go test ./... passes
