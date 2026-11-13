package smoke

import "testing"

func TestNewSet(t *testing.T) {
	intCases := []struct {
		name  string
		items []int
		want  int
	}{
		{"empty", nil, 0},
		{"unique", []int{1, 2, 3}, 3},
		{"duplicates", []int{1, 1, 2, 2}, 2},
	}
	for _, c := range intCases {
		t.Run("int/"+c.name, func(t *testing.T) {
			s := NewSet(c.items...)
			if got := SetLen(s); got != c.want {
				t.Errorf("len = %d, want %d", got, c.want)
			}
			for _, it := range c.items {
				if !SetHas(s, it) {
					t.Errorf("missing %d", it)
				}
			}
		})
	}
	strCases := []struct {
		name  string
		items []string
		want  int
	}{
		{"empty", nil, 0},
		{"unique", []string{"a", "b"}, 2},
		{"duplicates", []string{"a", "a", "b", "a"}, 2},
	}
	for _, c := range strCases {
		t.Run("string/"+c.name, func(t *testing.T) {
			s := NewSet(c.items...)
			if got := SetLen(s); got != c.want {
				t.Errorf("len = %d, want %d", got, c.want)
			}
			for _, it := range c.items {
				if !SetHas(s, it) {
					t.Errorf("missing %q", it)
				}
			}
		})
	}
}

func TestSetAdd(t *testing.T) {
	intCases := []struct {
		name    string
		initial []int
		add     int
		wantLen int
	}{
		{"new", []int{1, 2}, 3, 3},
		{"existing", []int{1, 2}, 2, 2},
		{"to empty", nil, 7, 1},
	}
	for _, c := range intCases {
		t.Run("int/"+c.name, func(t *testing.T) {
			s := NewSet(c.initial...)
			SetAdd(s, c.add)
			if !SetHas(s, c.add) || SetLen(s) != c.wantLen {
				t.Errorf("has=%v len=%d, want len %d", SetHas(s, c.add), SetLen(s), c.wantLen)
			}
		})
	}
	strCases := []struct {
		name    string
		initial []string
		add     string
		wantLen int
	}{
		{"new", []string{"a"}, "b", 2},
		{"existing", []string{"a"}, "a", 1},
		{"to empty", nil, "x", 1},
	}
	for _, c := range strCases {
		t.Run("string/"+c.name, func(t *testing.T) {
			s := NewSet(c.initial...)
			SetAdd(s, c.add)
			if !SetHas(s, c.add) || SetLen(s) != c.wantLen {
				t.Errorf("has=%v len=%d, want len %d", SetHas(s, c.add), SetLen(s), c.wantLen)
			}
		})
	}
}

func TestSetHas(t *testing.T) {
	intCases := []struct {
		name  string
		items []int
		probe int
		want  bool
	}{
		{"present", []int{1, 2, 3}, 2, true},
		{"absent", []int{1, 2, 3}, 4, false},
		{"empty", nil, 0, false},
	}
	for _, c := range intCases {
		t.Run("int/"+c.name, func(t *testing.T) {
			if got := SetHas(NewSet(c.items...), c.probe); got != c.want {
				t.Errorf("SetHas = %v, want %v", got, c.want)
			}
		})
	}
	strCases := []struct {
		name  string
		items []string
		probe string
		want  bool
	}{
		{"present", []string{"a", "b"}, "a", true},
		{"absent", []string{"a", "b"}, "c", false},
		{"empty", nil, "", false},
	}
	for _, c := range strCases {
		t.Run("string/"+c.name, func(t *testing.T) {
			if got := SetHas(NewSet(c.items...), c.probe); got != c.want {
				t.Errorf("SetHas = %v, want %v", got, c.want)
			}
		})
	}
}

func TestSetLen(t *testing.T) {
	intCases := []struct {
		name  string
		items []int
		want  int
	}{
		{"empty", nil, 0},
		{"three", []int{1, 2, 3}, 3},
		{"dups", []int{5, 5}, 1},
	}
	for _, c := range intCases {
		t.Run("int/"+c.name, func(t *testing.T) {
			if got := SetLen(NewSet(c.items...)); got != c.want {
				t.Errorf("SetLen = %d, want %d", got, c.want)
			}
		})
	}
	strCases := []struct {
		name  string
		items []string
		want  int
	}{
		{"empty", nil, 0},
		{"two", []string{"a", "b"}, 2},
		{"dups", []string{"a", "a"}, 1},
	}
	for _, c := range strCases {
		t.Run("string/"+c.name, func(t *testing.T) {
			if got := SetLen(NewSet(c.items...)); got != c.want {
				t.Errorf("SetLen = %d, want %d", got, c.want)
			}
		})
	}
}
