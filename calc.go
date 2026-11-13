// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Double returns a times 2.
func Double(a int) int { return a * 2 }

// Half returns a divided by 2, rounding toward zero.
func Half(a int) int { return a / 2 }

// Quad returns a times 4.
func Quad(a int) int { return a * 4 }

// Quarter returns a divided by 4, rounding toward zero.
func Quarter(a int) int { return a / 4 }
