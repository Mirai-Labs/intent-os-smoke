// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

import "sort"

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Negate returns minus a.
func Negate(a int) int { return -a }

// Double returns a times 2.
func Double(a int) int { return a * 2 }

// Half returns a divided by 2, rounding toward zero.
func Half(a int) int { return a / 2 }

// Quad returns a times 4.
func Quad(a int) int { return a * 4 }

// Quarter returns a divided by 4, rounding toward zero.
func Quarter(a int) int { return a / 4 }

// Clamp returns v limited to the range lo..hi.
func Clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// Median returns the median of xs: the middle value of the sorted values, or
// the mean of the two middle values for an even count. It returns 0 for an
// empty slice. The input slice is not modified.
func Median(xs []int) float64 {
	n := len(xs)
	if n == 0 {
		return 0
	}
	s := make([]int, n)
	copy(s, xs)
	sort.Ints(s)
	if n%2 == 1 {
		return float64(s[n/2])
	}
	return (float64(s[n/2-1]) + float64(s[n/2])) / 2
}
