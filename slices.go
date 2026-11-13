package smoke

// Map returns a new slice with f applied to every element of xs.
func Map[T, R any](xs []T, f func(T) R) []R {
	out := make([]R, 0, len(xs))
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

// Filter returns the elements of xs for which pred returns true, in order.
func Filter[T any](xs []T, pred func(T) bool) []T {
	out := make([]T, 0)
	for _, x := range xs {
		if pred(x) {
			out = append(out, x)
		}
	}
	return out
}

// Reduce folds xs into a single value, starting from init.
func Reduce[T, A any](xs []T, init A, f func(A, T) A) A {
	acc := init
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

// GroupBy groups the elements of xs by key, keeping input order inside each group.
func GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, x := range xs {
		k := key(x)
		out[k] = append(out[k], x)
	}
	return out
}

// Zip pairs elements of as and bs by index, stopping at the shorter slice.
func Zip[A, B any](as []A, bs []B) []struct {
	First  A
	Second B
} {
	n := min(len(as), len(bs))
	out := make([]struct {
		First  A
		Second B
	}, n)
	for i := 0; i < n; i++ {
		out[i].First = as[i]
		out[i].Second = bs[i]
	}
	return out
}
