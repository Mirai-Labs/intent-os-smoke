// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Double returns a times 2.
func Double(a int) int { return a * 2 }

// Half returns a divided by 2, rounding toward zero.
func Half(a int) int { return a / 2 }

// Triple returns a times 3.
func Triple(a int) int { return a * 3 }

// Third returns a divided by 3, rounding toward zero.
func Third(a int) int { return a / 3 }
