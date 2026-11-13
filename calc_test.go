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
	tests := []struct {
		name string
		a    int
		want int
	}{
		{"zero", 0, 0},
		{"one", 1, 3},
		{"negative", -3, -9},
		{"large", 14, 42},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Triple(tc.a); got != tc.want {
				t.Errorf("Triple(%d) = %d, want %d", tc.a, got, tc.want)
			}
		})
	}
}

func TestThird(t *testing.T) {
	tests := []struct {
		name string
		a    int
		want int
	}{
		{"zero", 0, 0},
		{"exact", 9, 3},
		{"truncates positive", 10, 3},
		{"small positive", 2, 0},
		{"small negative", -2, 0},
		{"truncates negative", -10, -3},
		{"exact negative", -9, -3},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Third(tc.a); got != tc.want {
				t.Errorf("Third(%d) = %d, want %d", tc.a, got, tc.want)
			}
		})
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
