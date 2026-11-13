# do not edit — managed by intent-os (intent: eb32959d-d593-4eab-8107-60ee53907dba, version: v1)
# Product Requirements Document

**Title:** Add Gcd and Lcm functions to the smoke library

## Description

Extend package smoke in calc.go with Gcd(a, b int) int and Lcm(a, b int) int, with table-driven tests in calc_test.go. Do not change Add.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- calc.go exports Gcd(a, b int) int that returns the greatest common divisor of a and b
- calc.go exports Lcm(a, b int) int that returns the least common multiple of a and b
- calc_test.go has table-driven tests for Gcd and Lcm and go test ./... passes

## Definition of Done

_No definition of done items defined._
