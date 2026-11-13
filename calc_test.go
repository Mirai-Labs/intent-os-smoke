package smoke

import (
	"slices"
	"testing"
)

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
	tests := []struct {
		name    string
		a, want int
	}{
		{"zero", 0, 0},
		{"one", 1, 4},
		{"negative", -3, -12},
		{"large", 21, 84},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Quad(tc.a); got != tc.want {
				t.Errorf("Quad(%d) = %d, want %d", tc.a, got, tc.want)
			}
		})
	}
}

func TestQuarter(t *testing.T) {
	tests := []struct {
		name    string
		a, want int
	}{
		{"zero", 0, 0},
		{"exact", 8, 2},
		{"positive truncates", 9, 2},
		{"small positive", 3, 0},
		{"small negative", -3, 0},
		{"negative truncates toward zero", -9, -2},
		{"negative exact", -8, -2},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Quarter(tc.a); got != tc.want {
				t.Errorf("Quarter(%d) = %d, want %d", tc.a, got, tc.want)
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

func TestMedian(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want float64
	}{
		{"empty", []int{}, 0},
		{"nil", nil, 0},
		{"single", []int{7}, 7},
		{"odd sorted", []int{1, 2, 3}, 2},
		{"even sorted", []int{1, 2, 3, 4}, 2.5},
		{"odd unsorted", []int{5, 1, 4, 2, 3}, 3},
		{"even unsorted", []int{9, 1, 7, 3}, 5},
		{"negatives", []int{-5, -1, -3, -2}, -2.5},
		{"duplicates", []int{2, 2, 2, 2}, 2},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Median(tc.in); got != tc.want {
				t.Errorf("Median(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestMedianDoesNotModifyInput(t *testing.T) {
	in := []int{5, 3, 9, 1, 7, 2}
	before := slices.Clone(in)
	Median(in)
	if !slices.Equal(in, before) {
		t.Errorf("Median modified input: got %v, want %v", in, before)
	}
}
