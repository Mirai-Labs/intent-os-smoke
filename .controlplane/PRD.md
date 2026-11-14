# do not edit — managed by intent-os (intent: 8effa0c8-7ee1-4d3f-a8c7-846298d5ef10, version: v1)
# Product Requirements Document

**Title:** Add Abs to the smoke library

## Description

Add Abs(a int) int to calc.go returning the absolute value, with table-driven tests in calc_test.go. Do not change existing functions.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Abs(a int) int that returns the absolute value of a
- Abs(-3) returns 3 and Abs(3) returns 3 and Abs(0) returns 0
- calc_test.go has table-driven tests for Abs and go test ./... passes

## Definition of Done

_No definition of done items defined._
