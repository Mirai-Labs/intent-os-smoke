# do not edit — managed by intent-os (intent: 29a2941a-652e-4224-8b42-e9d8a99cc248, version: v1)
# Acceptance Criteria

- calc.go exports Sub(a, b int) int that returns a minus b
- calc.go exports Mul(a, b int) int that returns a times b
- calc_test.go has table-driven tests for Sub and Mul and go test ./... passes
