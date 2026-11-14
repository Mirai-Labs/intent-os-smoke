# do not edit — managed by intent-os (intent: 12a64873-84f3-4161-b24b-eacd2ba87d41, version: v1)
# Task Checklist — Add Clamp to the smoke library

## Acceptance Criteria Tasks

- [ ] calc.go exports Clamp(v, lo, hi int) int that returns lo when v<lo, hi when v>hi and v otherwise
- [ ] Clamp(5,0,3) returns 3, Clamp(-1,0,3) returns 0 and Clamp(2,0,3) returns 2
- [ ] calc_test.go has table-driven tests for Clamp and go test ./... passes

## Definition of Done Tasks

_No definition of done tasks defined._
