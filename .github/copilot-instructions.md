# do not edit — managed by intent-os (intent: a90051f4-607b-446f-8f4e-8996352c435e, version: v1)
# GitHub Copilot Instructions — Add Max and Min functions to the smoke library

> This file is managed by intent-os. Do not edit manually.

## Intent Summary

Extend package smoke in calc.go with Max(a, b int) int and Min(a, b int) int, with table-driven tests in calc_test.go. Do not change Add.

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
