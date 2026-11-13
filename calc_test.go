package smoke

import "testing"

func TestDouble(t *testing.T) {
	cases := []struct{ a, want int }{{0, 0}, {1, 2}, {-3, -6}, {21, 42}}
	for _, c := range cases {
		if got := Double(c.a); got != c.want {
			t.Errorf("Double(%d) = %d, want %d", c.a, got, c.want)
		}
	}
}

func TestHalf(t *testing.T) {
	cases := []struct{ a, want int }{{0, 0}, {4, 2}, {5, 2}, {1, 0}, {-1, 0}, {-5, -2}, {-4, -2}}
	for _, c := range cases {
		if got := Half(c.a); got != c.want {
			t.Errorf("Half(%d) = %d, want %d", c.a, got, c.want)
		}
	}
}

func TestTriple(t *testing.T) {
	cases := []struct{ a, want int }{{0, 0}, {1, 3}, {-3, -9}, {14, 42}}
	for _, c := range cases {
		if got := Triple(c.a); got != c.want {
			t.Errorf("Triple(%d) = %d, want %d", c.a, got, c.want)
		}
	}
}

func TestThird(t *testing.T) {
	cases := []struct{ a, want int }{{0, 0}, {9, 3}, {10, 3}, {2, 0}, {-2, 0}, {-10, -3}, {-9, -3}}
	for _, c := range cases {
		if got := Third(c.a); got != c.want {
			t.Errorf("Third(%d) = %d, want %d", c.a, got, c.want)
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
