# do not edit — managed by intent-os (intent: 12a64873-84f3-4161-b24b-eacd2ba87d41, version: v1)
# Product Requirements Document

**Title:** Add Clamp to the smoke library

## Description

Add Clamp(v, lo, hi int) int to calc.go returning v limited to the range lo..hi, with table-driven tests in calc_test.go. Do not change existing functions.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Clamp(v, lo, hi int) int that returns lo when v<lo, hi when v>hi and v otherwise
- Clamp(5,0,3) returns 3, Clamp(-1,0,3) returns 0 and Clamp(2,0,3) returns 2
- calc_test.go has table-driven tests for Clamp and go test ./... passes

## Definition of Done

_No definition of done items defined._
