package smoke

import "testing"

func TestMax(t *testing.T) {
	cases := []struct{ a, b, want int }{{1, 2, 2}, {2, 1, 2}, {3, 3, 3}, {-1, -2, -1}, {0, -5, 0}}
	for _, c := range cases {
		if got := Max(c.a, c.b); got != c.want {
			t.Errorf("Max(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestMin(t *testing.T) {
	cases := []struct{ a, b, want int }{{1, 2, 1}, {2, 1, 1}, {3, 3, 3}, {-1, -2, -2}, {0, -5, -5}}
	for _, c := range cases {
		if got := Min(c.a, c.b); got != c.want {
			t.Errorf("Min(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
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
