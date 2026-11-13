# do not edit — managed by intent-os (intent: 93ad5f83-66cd-4b2c-8b56-122a45890c58, version: v1)
# Task Checklist — Add Median to the smoke library

## Acceptance Criteria Tasks

- [ ] calc.go exports Median(xs []int) float64: the middle value of the sorted values, the mean of the two middle values for an even count, and 0 for an empty slice
- [ ] Median never modifies or reorders the slice passed by the caller, and calc_test.go proves it with a test that compares the input before and after
- [ ] calc_test.go has table-driven tests for odd, even, empty and unsorted inputs, and go test ./... passes

## Definition of Done Tasks

_No definition of done tasks defined._
