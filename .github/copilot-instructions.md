# do not edit — managed by intent-os (intent: 8effa0c8-7ee1-4d3f-a8c7-846298d5ef10, version: v1)
# GitHub Copilot Instructions — Add Abs to the smoke library

> This file is managed by intent-os. Do not edit manually.

## Intent Summary

Add Abs(a int) int to calc.go returning the absolute value, with table-driven tests in calc_test.go. Do not change existing functions.

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
