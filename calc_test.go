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

func TestQuad(t *testing.T) {
	cases := []struct{ a, want int }{{0, 0}, {1, 4}, {-3, -12}, {21, 84}}
	for _, c := range cases {
		if got := Quad(c.a); got != c.want {
			t.Errorf("Quad(%d) = %d, want %d", c.a, got, c.want)
		}
	}
}

func TestQuarter(t *testing.T) {
	cases := []struct{ a, want int }{{0, 0}, {8, 2}, {9, 2}, {3, 0}, {-3, 0}, {-9, -2}, {-8, -2}}
	for _, c := range cases {
		if got := Quarter(c.a); got != c.want {
			t.Errorf("Quarter(%d) = %d, want %d", c.a, got, c.want)
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
