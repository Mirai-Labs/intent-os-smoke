package smoke

import "testing"

func TestSub(t *testing.T) {
	cases := []struct{ a, b, want int }{{3, 2, 1}, {0, 0, 0}, {1, 3, -2}, {-1, -1, 0}}
	for _, c := range cases {
		if got := Sub(c.a, c.b); got != c.want {
			t.Errorf("Sub(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestMul(t *testing.T) {
	cases := []struct{ a, b, want int }{{2, 3, 6}, {0, 5, 0}, {-2, 3, -6}, {-2, -3, 6}}
	for _, c := range cases {
		if got := Mul(c.a, c.b); got != c.want {
			t.Errorf("Mul(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
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
