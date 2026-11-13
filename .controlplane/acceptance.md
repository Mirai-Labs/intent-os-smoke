# do not edit — managed by intent-os (intent: 3ebdfbba-af20-426e-b0a3-984e8f71abe2, version: v1)
# Acceptance Criteria

- calc.go exports Quad(a int) int that returns a times 4
- calc.go exports Quarter(a int) int that returns a divided by 4, rounding toward zero
- calc_test.go has table-driven tests for Quad and Quarter and go test ./... passes
