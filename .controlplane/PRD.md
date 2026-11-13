# do not edit — managed by intent-os (intent: edd6c7ad-7b59-46ee-875f-b42dccc61b36, version: v1)
# Product Requirements Document

**Title:** Add a generic Set type to the smoke library

## Description

Add a new file set.go to package smoke that declares the generic type alias `type Set[T comparable] = map[T]struct{}` (a generic type alias) plus the generic functions NewSet, SetAdd, SetHas and SetLen over it, and set_test.go with table-driven tests. Do not change calc.go or calc_test.go.

## Functional Requirements

_No functional requirements defined._

## Non-Functional Requirements

_No non-functional requirements defined._

## Acceptance Criteria

- set.go declares `type Set[T comparable] = map[T]struct{}` as a generic type alias (not a defined type) and NewSet[T comparable](items ...T) Set[T]
- set.go exports SetAdd, SetHas and SetLen as generic functions over Set[T]
- set_test.go has table-driven tests for NewSet, SetAdd, SetHas and SetLen with int and string element types, and go test ./... passes

## Definition of Done

_No definition of done items defined._
