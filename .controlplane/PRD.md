# do not edit — managed by intent-os (intent: 0eae61fb-0d99-4634-a0d9-dc92326a7187, version: v1)
# Product Requirements Document

**Title:** Add Negate and a CI workflow to the smoke library

## Description

Add Negate(a int) int to calc.go with table-driven tests in calc_test.go. Do not change Add. Also add a GitHub Actions workflow at .github/workflows/test.yml that runs go test ./... on every push.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Negate(a int) int that returns minus a
- Negate(0) returns 0 and Negate(Negate(x)) returns x
- calc_test.go has table-driven tests for Negate and go test ./... passes

## Definition of Done

_No definition of done items defined._
