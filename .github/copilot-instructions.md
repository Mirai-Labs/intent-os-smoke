# do not edit — managed by intent-os (intent: edd6c7ad-7b59-46ee-875f-b42dccc61b36, version: v1)
# GitHub Copilot Instructions — Add a generic Set type to the smoke library

> This file is managed by intent-os. Do not edit manually.

## Intent Summary

Add a new file set.go to package smoke that declares the generic type alias `type Set[T comparable] = map[T]struct{}` (a generic type alias) plus the generic functions NewSet, SetAdd, SetHas and SetLen over it, and set_test.go with table-driven tests. Do not change calc.go or calc_test.go.

## Control-Plane Artifacts

This repository is governed by an intent-os harness. Key artifacts:

- `AGENTS.md — vendor-neutral agent instructions (read this first)`
- `.controlplane/manifest.yaml — list of all managed artifacts`
- `.controlplane/constraints.md — allowed/denied tools and active policies`
- `.controlplane/acceptance.md — acceptance criteria and definition of done`
- `.controlplane/PRD.md — product requirements document`
- `.controlplane/TASK.md — actionable task checklist`
- `.controlplane/guardrails.json — machine-readable tool guardrails`
- `.controlplane/context-map.json — retrieved code/symbol/doc context`

## Key Constraints

Allowed tools: `read`, `write`, `grep`, `run`

Denied tools: `network_egress`

Do not modify files listed in `.controlplane/manifest.yaml` — they are regenerated on every bundle pull.
