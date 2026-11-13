# do not edit — managed by intent-os (intent: b3800831-8889-429e-b4a7-f50f25327ee3, version: v1)
# Product Requirements Document

**Title:** Add IsEven and IsOdd functions to the smoke library

## Description

Extend package smoke in calc.go with IsEven(a int) bool and IsOdd(a int) bool, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports IsEven(a int) bool that reports whether a is even
- calc.go exports IsOdd(a int) bool that reports whether a is odd
- calc_test.go has table-driven tests for IsEven and IsOdd and go test ./... passes

## Definition of Done

_No definition of done items defined._
