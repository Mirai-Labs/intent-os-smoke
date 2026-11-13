# do not edit — managed by intent-os (intent: a90051f4-607b-446f-8f4e-8996352c435e, version: v1)
# Product Requirements Document

**Title:** Add Max and Min functions to the smoke library

## Description

Extend package smoke in calc.go with Max(a, b int) int and Min(a, b int) int, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Max(a, b int) int that returns the larger of a and b
- calc.go exports Min(a, b int) int that returns the smaller of a and b
- calc_test.go has table-driven tests for Max and Min and go test ./... passes

## Definition of Done

_No definition of done items defined._
