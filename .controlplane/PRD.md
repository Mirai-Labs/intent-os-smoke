# do not edit — managed by intent-os (intent: 93ad5f83-66cd-4b2c-8b56-122a45890c58, version: v1)
# Product Requirements Document

**Title:** Add Median to the smoke library

## Description

Extend package smoke in calc.go with Median(xs []int) float64, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Median(xs []int) float64: the middle value of the sorted values, the mean of the two middle values for an even count, and 0 for an empty slice
- Median never modifies or reorders the slice passed by the caller, and calc_test.go proves it with a test that compares the input before and after
- calc_test.go has table-driven tests for odd, even, empty and unsorted inputs, and go test ./... passes

## Definition of Done

_No definition of done items defined._
