# do not edit — managed by intent-os (intent: 4093edd6-8246-4fe2-9d25-8ca8dda5ee85, version: v1)
# Acceptance Criteria

- calc.go exports Clamp(v, lo, hi int) int that limits v to the range lo..hi
- calc.go exports Sign(a int) int that returns -1, 0 or 1 according to the sign of a
- calc_test.go has table-driven tests for Clamp and Sign and go test ./... passes
