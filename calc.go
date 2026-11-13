// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// IsEven reports whether a is even.
func IsEven(a int) bool { return a%2 == 0 }

// IsOdd reports whether a is odd.
func IsOdd(a int) bool { return a%2 != 0 }
