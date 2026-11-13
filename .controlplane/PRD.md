# do not edit — managed by intent-os (intent: 29a2c343-5613-4412-81ef-5a9598fb3be8, version: v1)
# Product Requirements Document

**Title:** Add Square and Cube functions to the smoke library

## Description

Extend package smoke in calc.go with Square(a int) int and Cube(a int) int, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Square(a int) int that returns a times a
- calc.go exports Cube(a int) int that returns a times a times a
- calc_test.go has table-driven tests for Square and Cube and go test ./... passes

## Definition of Done

_No definition of done items defined._
