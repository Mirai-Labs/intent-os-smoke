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

var parityCases = []struct {
	a    int
	even bool
}{{0, true}, {1, false}, {2, true}, {-1, false}, {-2, true}, {7, false}, {10, true}}

func TestIsEven(t *testing.T) {
	for _, c := range parityCases {
		if got := IsEven(c.a); got != c.even {
			t.Errorf("IsEven(%d) = %v, want %v", c.a, got, c.even)
		}
	}
}

func TestIsOdd(t *testing.T) {
	for _, c := range parityCases {
		if got := IsOdd(c.a); got != !c.even {
			t.Errorf("IsOdd(%d) = %v, want %v", c.a, got, !c.even)
		}
	}
}
