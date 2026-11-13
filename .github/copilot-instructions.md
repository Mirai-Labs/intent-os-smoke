# do not edit — managed by intent-os (intent: 0a5bc799-9899-46f6-9fd5-f7c0cfd03e6e, version: v1)
# GitHub Copilot Instructions — Add Abs and Neg functions to the smoke library

> This file is managed by intent-os. Do not edit manually.

## Intent Summary

Extend package smoke in calc.go with Abs(a int) int and Neg(a int) int, with table-driven tests in calc_test.go. Do not change Add.

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
