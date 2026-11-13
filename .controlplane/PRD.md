# do not edit — managed by intent-os (intent: cc760661-70af-42e9-b582-7131f60dbfbf, version: v1)
# Product Requirements Document

**Title:** Add Double and Half functions to the smoke library

## Description

Extend package smoke in calc.go with Double(a int) int and Half(a int) int, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Double(a int) int that returns a times 2
- calc.go exports Half(a int) int that returns a divided by 2, rounding toward zero
- calc_test.go has table-driven tests for Double and Half and go test ./... passes

## Definition of Done

_No definition of done items defined._
