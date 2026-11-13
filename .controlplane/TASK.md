# do not edit — managed by intent-os (intent: edd6c7ad-7b59-46ee-875f-b42dccc61b36, version: v1)
# Task Checklist — Add a generic Set type to the smoke library

## Acceptance Criteria Tasks

- [ ] set.go declares `type Set[T comparable] = map[T]struct{}` as a generic type alias (not a defined type) and NewSet[T comparable](items ...T) Set[T]
- [ ] set.go exports SetAdd, SetHas and SetLen as generic functions over Set[T]
- [ ] set_test.go has table-driven tests for NewSet, SetAdd, SetHas and SetLen with int and string element types, and go test ./... passes

## Definition of Done Tasks

_No definition of done tasks defined._
