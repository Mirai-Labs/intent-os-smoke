package smoke

import "testing"

func TestSquare(t *testing.T) {
	cases := []struct{ a, want int }{{0, 0}, {1, 1}, {3, 9}, {-4, 16}}
	for _, c := range cases {
		if got := Square(c.a); got != c.want {
			t.Errorf("Square(%d) = %d, want %d", c.a, got, c.want)
		}
	}
}

func TestCube(t *testing.T) {
	cases := []struct{ a, want int }{{0, 0}, {1, 1}, {3, 27}, {-2, -8}}
	for _, c := range cases {
		if got := Cube(c.a); got != c.want {
			t.Errorf("Cube(%d) = %d, want %d", c.a, got, c.want)
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
