# do not edit — managed by intent-os (intent: 0a9150c9-b275-435e-af86-4668c8909dcc, version: v1)
# GitHub Copilot Instructions — Change Add to also count one extra unit

> This file is managed by intent-os. Do not edit manually.

## Intent Summary

Change Add in calc.go so that Add(a, b) returns a plus b plus 1. Do not edit, delete or skip any existing test or test file.

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
