# do not edit — managed by intent-os (intent: 3865c6f8-8222-448b-a84e-d809ad9d9f72, version: v1)
# Product Requirements Document

**Title:** Add Triple and Third functions to the smoke library

## Description

Extend package smoke in calc.go with Triple(a int) int and Third(a int) int, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Triple(a int) int that returns a times 3
- calc.go exports Third(a int) int that returns a divided by 3, rounding toward zero
- calc_test.go has table-driven tests for Triple and Third and go test ./... passes

## Definition of Done

_No definition of done items defined._
