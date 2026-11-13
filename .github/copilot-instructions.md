# do not edit — managed by intent-os (intent: eb32959d-d593-4eab-8107-60ee53907dba, version: v1)
# GitHub Copilot Instructions — Add Gcd and Lcm functions to the smoke library

> This file is managed by intent-os. Do not edit manually.

## Intent Summary

Extend package smoke in calc.go with Gcd(a, b int) int and Lcm(a, b int) int, with table-driven tests in calc_test.go. Do not change Add.

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
