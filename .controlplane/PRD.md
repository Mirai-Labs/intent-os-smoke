# do not edit — managed by intent-os (intent: 29a2941a-652e-4224-8b42-e9d8a99cc248, version: v1)
# Product Requirements Document

**Title:** Add Sub and Mul functions to the smoke library

## Description

Extend package smoke in calc.go with Sub(a, b int) int and Mul(a, b int) int, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Sub(a, b int) int that returns a minus b
- calc.go exports Mul(a, b int) int that returns a times b
- calc_test.go has table-driven tests for Sub and Mul and go test ./... passes

## Definition of Done

_No definition of done items defined._
