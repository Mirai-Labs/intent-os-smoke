# do not edit — managed by intent-os (intent: 3865c6f8-8222-448b-a84e-d809ad9d9f72, version: v1)
# Acceptance Criteria

- calc.go exports Triple(a int) int that returns a times 3
- calc.go exports Third(a int) int that returns a divided by 3, rounding toward zero
- calc_test.go has table-driven tests for Triple and Third and go test ./... passes
