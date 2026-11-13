# do not edit — managed by intent-os (intent: cc760661-70af-42e9-b582-7131f60dbfbf, version: v1)
# Acceptance Criteria

- calc.go exports Double(a int) int that returns a times 2
- calc.go exports Half(a int) int that returns a divided by 2, rounding toward zero
- calc_test.go has table-driven tests for Double and Half and go test ./... passes
