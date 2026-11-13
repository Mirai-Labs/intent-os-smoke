// Package smoke is a tiny library used to exercise the Intent-OS delivery loop.
package smoke

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Gcd returns the greatest common divisor of a and b. The result is always
// non-negative, and Gcd(0, 0) is 0.
func Gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Lcm returns the least common multiple of a and b. The result is always
// non-negative, and is 0 if either argument is 0.
func Lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	l := a / Gcd(a, b) * b
	if l < 0 {
		l = -l
	}
	return l
}
