# do not edit — managed by intent-os (intent: 5f6aa215-1fef-429e-89e7-ac811c84b332, version: v1)
# Acceptance Criteria

- slices.go exports Map[T, R any](xs []T, f func(T) R) []R, Filter[T any](xs []T, pred func(T) bool) []T and Reduce[T, A any](xs []T, init A, f func(A, T) A) A
- slices.go exports GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T that keeps the input order inside each group, and Zip[A, B any](as []A, bs []B) []struct{ First A; Second B } that stops at the shorter slice
- slices_test.go has table-driven tests for Map, Filter, Reduce, GroupBy and Zip using at least two different type parameters each, and go test ./... passes
