// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Abs returns the absolute value of a.
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Neg returns minus a.
func Neg(a int) int { return -a }
