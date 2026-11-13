package smoke

import "testing"

func TestAbs(t *testing.T) {
	cases := []struct{ a, want int }{{5, 5}, {0, 0}, {-5, 5}}
	for _, c := range cases {
		if got := Abs(c.a); got != c.want {
			t.Errorf("Abs(%d) = %d, want %d", c.a, got, c.want)
		}
	}
}

func TestNeg(t *testing.T) {
	cases := []struct{ a, want int }{{5, -5}, {0, 0}, {-5, 5}}
	for _, c := range cases {
		if got := Neg(c.a); got != c.want {
			t.Errorf("Neg(%d) = %d, want %d", c.a, got, c.want)
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct{ a, b, want int }{{1, 2, 3}, {0, 0, 0}, {-1, 1, 0}}
	for _, c := range cases {
		if got := Add(c.a, c.b); got != c.want {
			t.Errorf("Add(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
