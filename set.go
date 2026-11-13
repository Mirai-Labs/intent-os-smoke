package smoke

// Set is a generic set of comparable elements.
type Set[T comparable] = map[T]struct{}

// NewSet returns a Set containing items.
func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, it := range items {
		s[it] = struct{}{}
	}
	return s
}

// SetAdd adds item to s.
func SetAdd[T comparable](s Set[T], item T) { s[item] = struct{}{} }

// SetHas reports whether item is in s.
func SetHas[T comparable](s Set[T], item T) bool {
	_, ok := s[item]
	return ok
}

// SetLen returns the number of elements in s.
func SetLen[T comparable](s Set[T]) int { return len(s) }
