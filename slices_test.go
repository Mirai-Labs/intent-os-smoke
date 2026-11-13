package smoke

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		cases := []struct {
			in   []int
			want []string
		}{
			{nil, []string{}},
			{[]int{1, 2, 3}, []string{"1", "2", "3"}},
		}
		for _, c := range cases {
			if got := Map(c.in, strconv.Itoa); !reflect.DeepEqual(got, c.want) {
				t.Errorf("Map(%v) = %v, want %v", c.in, got, c.want)
			}
		}
	})
	t.Run("string to int", func(t *testing.T) {
		cases := []struct {
			in   []string
			want []int
		}{
			{[]string{}, []int{}},
			{[]string{"a", "bcd", ""}, []int{1, 3, 0}},
		}
		for _, c := range cases {
			if got := Map(c.in, func(s string) int { return len(s) }); !reflect.DeepEqual(got, c.want) {
				t.Errorf("Map(%v) = %v, want %v", c.in, got, c.want)
			}
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		cases := []struct {
			in, want []int
		}{
			{nil, []int{}},
			{[]int{1, 2, 3, 4}, []int{2, 4}},
			{[]int{1, 3}, []int{}},
		}
		for _, c := range cases {
			if got := Filter(c.in, func(i int) bool { return i%2 == 0 }); !reflect.DeepEqual(got, c.want) {
				t.Errorf("Filter(%v) = %v, want %v", c.in, got, c.want)
			}
		}
	})
	t.Run("string", func(t *testing.T) {
		cases := []struct {
			in, want []string
		}{
			{[]string{"go", "rust", "gleam"}, []string{"go", "gleam"}},
			{[]string{}, []string{}},
		}
		for _, c := range cases {
			if got := Filter(c.in, func(s string) bool { return strings.HasPrefix(s, "g") }); !reflect.DeepEqual(got, c.want) {
				t.Errorf("Filter(%v) = %v, want %v", c.in, got, c.want)
			}
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("int sum", func(t *testing.T) {
		cases := []struct {
			in   []int
			init int
			want int
		}{
			{nil, 7, 7},
			{[]int{1, 2, 3}, 0, 6},
			{[]int{1, 2, 3}, 10, 16},
		}
		for _, c := range cases {
			if got := Reduce(c.in, c.init, func(a, x int) int { return a + x }); got != c.want {
				t.Errorf("Reduce(%v, %d) = %d, want %d", c.in, c.init, got, c.want)
			}
		}
	})
	t.Run("string to string", func(t *testing.T) {
		cases := []struct {
			in   []string
			init string
			want string
		}{
			{nil, "x", "x"},
			{[]string{"a", "b", "c"}, "", "abc"},
		}
		for _, c := range cases {
			if got := Reduce(c.in, c.init, func(a string, x string) string { return a + x }); got != c.want {
				t.Errorf("Reduce(%v, %q) = %q, want %q", c.in, c.init, got, c.want)
			}
		}
	})
	t.Run("string to int", func(t *testing.T) {
		got := Reduce([]string{"a", "bc", "def"}, 0, func(a int, s string) int { return a + len(s) })
		if got != 6 {
			t.Errorf("got %d, want 6", got)
		}
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("int by parity", func(t *testing.T) {
		cases := []struct {
			in   []int
			want map[bool][]int
		}{
			{nil, map[bool][]int{}},
			{[]int{1, 2, 3, 4, 5}, map[bool][]int{true: {2, 4}, false: {1, 3, 5}}},
		}
		for _, c := range cases {
			if got := GroupBy(c.in, func(i int) bool { return i%2 == 0 }); !reflect.DeepEqual(got, c.want) {
				t.Errorf("GroupBy(%v) = %v, want %v", c.in, got, c.want)
			}
		}
	})
	t.Run("string by length", func(t *testing.T) {
		cases := []struct {
			in   []string
			want map[int][]string
		}{
			{[]string{"b", "aa", "a", "bb", "c"}, map[int][]string{1: {"b", "a", "c"}, 2: {"aa", "bb"}}},
			{[]string{}, map[int][]string{}},
		}
		for _, c := range cases {
			if got := GroupBy(c.in, func(s string) int { return len(s) }); !reflect.DeepEqual(got, c.want) {
				t.Errorf("GroupBy(%v) = %v, want %v", c.in, got, c.want)
			}
		}
	})
}

func TestZip(t *testing.T) {
	type pair = struct {
		First  int
		Second string
	}
	t.Run("int string", func(t *testing.T) {
		cases := []struct {
			as   []int
			bs   []string
			want []pair
		}{
			{nil, nil, []pair{}},
			{[]int{1, 2, 3}, []string{"a", "b", "c"}, []pair{{1, "a"}, {2, "b"}, {3, "c"}}},
			{[]int{1, 2, 3}, []string{"a"}, []pair{{1, "a"}}},
			{[]int{1}, []string{"a", "b"}, []pair{{1, "a"}}},
		}
		for _, c := range cases {
			if got := Zip(c.as, c.bs); !reflect.DeepEqual(got, c.want) {
				t.Errorf("Zip(%v, %v) = %v, want %v", c.as, c.bs, got, c.want)
			}
		}
	})
	t.Run("bool float", func(t *testing.T) {
		type pair2 = struct {
			First  bool
			Second float64
		}
		cases := []struct {
			as   []bool
			bs   []float64
			want []pair2
		}{
			{[]bool{true, false}, []float64{1.5}, []pair2{{true, 1.5}}},
			{[]bool{true}, []float64{}, []pair2{}},
		}
		for _, c := range cases {
			if got := Zip(c.as, c.bs); !reflect.DeepEqual(got, c.want) {
				t.Errorf("Zip(%v, %v) = %v, want %v", c.as, c.bs, got, c.want)
			}
		}
	})
}
