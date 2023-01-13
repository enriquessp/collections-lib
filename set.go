package collections

// A collection of unique comparable items. Uses a map with only true values
// to accomplish set functionality.
type Set[T comparable] map[T]bool

// Create a new empty set with the specified initial size.
func NewSet[T comparable](size int) Set[T] {
	return make(Set[T], size)
}

// Add a new key to the set
func (s Set[T]) Add(key T) {
	s[key] = true
}

// Remove a key from the set. If the key is not in the set then noop
func (s Set[T]) Remove(key T) {
	delete(s, key)
}

// Check if Set s contains key
func (s Set[T]) Contains(key T) bool {
	return s[key]
}

// A union B
func (a Set[T]) Union(b Set[T]) Set[T] {
	small, large := smallLarge(a, b)

	for key := range small {
		large.Add(key)
	}
	return large
}

// A intersect B
func (a Set[T]) Intersection(b Set[T]) Set[T] {
	small, large := smallLarge(a, b)

	resultSet := NewSet[T](0)
	for key := range small {
		if large.Contains(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

// A compliment
func (a Set[T]) Complement(b Set[T]) Set[T] {
	resultSet := NewSet[T](0)
	for key := range b {
		if !a.Contains(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

// A difference B | NOTE: A-B != B-A
func (a Set[T]) Difference(b Set[T]) Set[T] {
	resultSet := NewSet[T](0)
	for key := range a {
		if !b.Contains(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

// A == B (all elements of A are in B and vice versa)
func (a Set[T]) Equals(b Set[T]) bool {
	return len(a.Difference(b)) == 0 && len(b.Difference(a)) == 0
}

// Turn a Set into a slice
func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}

	return slice
}

// Filter Set elements
func (s Set[T]) Filter(predicate func(t T) bool) Set[T] {
	rs := NewSet[T](0)

	for e := range s {
		if predicate(e) {
			rs.Add(e)
		}
	}

	return rs
}

// Map convert Set elements to another type
func Map[S, T comparable](set Set[T], convert func(t T) S) Set[S] {
	rs := NewSet[S](0)

	for e := range set {
		rs.Add(convert(e))
	}

	return rs
}

// returns the small and large sets according to their len
func smallLarge[T comparable](a, b Set[T]) (Set[T], Set[T]) {
	small, large := b, a
	if len(b) > len(a) {
		small, large = a, b
	}

	return small, large
}
