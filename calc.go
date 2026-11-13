// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Square returns a multiplied by itself.
func Square(a int) int { return a * a }

// Cube returns a multiplied by itself twice.
func Cube(a int) int { return a * a * a }
