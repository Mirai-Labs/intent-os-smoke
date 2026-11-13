package smoke

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct{ a, b, want int }{{1, 2, 3}, {0, 0, 0}, {-1, 1, 0}}
	for _, c := range cases {
		if got := Add(c.a, c.b); got != c.want {
			t.Errorf("Add(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestGcd(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{12, 18, 6}, {18, 12, 6}, {7, 13, 1}, {0, 5, 5}, {5, 0, 5},
		{0, 0, 0}, {-12, 18, 6}, {12, -18, 6}, {-12, -18, 6}, {10, 10, 10},
	}
	for _, c := range cases {
		if got := Gcd(c.a, c.b); got != c.want {
			t.Errorf("Gcd(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestLcm(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{4, 6, 12}, {6, 4, 12}, {7, 13, 91}, {0, 5, 0}, {5, 0, 0},
		{0, 0, 0}, {-4, 6, 12}, {4, -6, 12}, {-4, -6, 12}, {10, 10, 10},
	}
	for _, c := range cases {
		if got := Lcm(c.a, c.b); got != c.want {
			t.Errorf("Lcm(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
