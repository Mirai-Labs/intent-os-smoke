// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Clamp limits v to the range lo..hi. If lo > hi, the bounds are swapped.
func Clamp(v, lo, hi int) int {
	if lo > hi {
		lo, hi = hi, lo
	}
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// Sign returns -1, 0 or 1 according to the sign of a.
func Sign(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	}
	return 0
}
