# do not edit — managed by intent-os (intent: 0a5bc799-9899-46f6-9fd5-f7c0cfd03e6e, version: v1)
# Product Requirements Document

**Title:** Add Abs and Neg functions to the smoke library

## Description

Extend package smoke in calc.go with Abs(a int) int and Neg(a int) int, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Abs(a int) int that returns the absolute value of a
- calc.go exports Neg(a int) int that returns minus a
- calc_test.go has table-driven tests for Abs and Neg and go test ./... passes

## Definition of Done

_No definition of done items defined._
