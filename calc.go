// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

// Add returns the sum of a and b plus one.
func Add(a, b int) int { return a + b + 1 }

// Double returns a times 2.
func Double(a int) int { return a * 2 }

// Half returns a divided by 2, rounding toward zero.
func Half(a int) int { return a / 2 }
