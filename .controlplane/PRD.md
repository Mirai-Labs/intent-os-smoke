# do not edit — managed by intent-os (intent: 3ebdfbba-af20-426e-b0a3-984e8f71abe2, version: v1)
# Product Requirements Document

**Title:** Add Quad and Quarter functions to the smoke library

## Description

Extend package smoke in calc.go with Quad(a int) int and Quarter(a int) int, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Quad(a int) int that returns a times 4
- calc.go exports Quarter(a int) int that returns a divided by 4, rounding toward zero
- calc_test.go has table-driven tests for Quad and Quarter and go test ./... passes

## Definition of Done

_No definition of done items defined._
