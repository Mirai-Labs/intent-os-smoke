# do not edit — managed by intent-os (intent: 5f6aa215-1fef-429e-89e7-ac811c84b332, version: v1)
# GitHub Copilot Instructions — Add generic slice helpers (second attempt)

> This file is managed by intent-os. Do not edit manually.

## Intent Summary

Add a new file slices.go to package smoke with the generic helpers Map, Filter, Reduce, GroupBy and Zip, and slices_test.go with table-driven tests for each of them. Do not change calc.go or calc_test.go.

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
