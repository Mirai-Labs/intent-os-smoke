# do not edit — managed by intent-os (intent: 0a9150c9-b275-435e-af86-4668c8909dcc, version: v1)
# AGENTS.md — Change Add to also count one extra unit

> Vendor-neutral agent instructions managed by intent-os.

## Scope

Change Add in calc.go so that Add(a, b) returns a plus b plus 1. Do not edit, delete or skip any existing test or test file.

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

- Add(a, b int) int in calc.go returns a+b+1
- no existing test file is modified, deleted or skipped
- go test ./... passes
