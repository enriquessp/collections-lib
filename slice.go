package collections

// -------------------------------------------------
// SLICE HELPERS

// Create a Set from a slice.
func SliceToSet[T comparable](s []T) Set[T] {
	set := NewSet[T](len(s))
	for _, item := range s {
		set.Add(item)
	}
	return set
}

func GroupedBySlice[T any, K comparable](s []T, key func(s T) K) map[K][]T {
	rs := make(map[K][]T)

	for _, v := range s {
		rs[key(v)] = append(rs[key(v)], v)
	}

	return rs
}

// Foreach
func ForeachSlice[T any](s []T, apply func(t T)) {
	for _, e := range s {
		apply(e)
	}
}

// Map a slice to a set using a function f
func MapSliceToSet[S any, T comparable](s []S, f func(s S) T) Set[T] {
	set := NewSet[T](len(s))
	for _, item := range s {
		set.Add(f(item))
	}
	return set
}

// Union two slices. The provided slices do not need to be unique. Order not guaranteed.
func SliceUnion[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	union := aSet.Union(bSet)
	return union.ToSlice()
}

// Intersection of two slices. The provided slices do not need to be unique. Order not guaranteed.
func SliceIntersection[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	intersection := aSet.Intersection(bSet)
	return intersection.ToSlice()
}

// Complement of A with regards to B. Slices do not need to be unique. Order not guaranteed.
func SliceComplement[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	complement := aSet.Complement(bSet)
	return complement.ToSlice()
}

// Difference of A-B. Slices do not need to be unique. Order not guaranteed.
func SliceDifference[T comparable](a, b []T) []T {
	aSet, bSet := SliceToSet(a), SliceToSet(b)
	difference := aSet.Difference(bSet)
	return difference.ToSlice()
}

// Filter slice elements
func FilterSlice[T any](s []T, predicate func(t T) bool) []T {
	rs := []T{}

	for _, e := range s {
		if predicate(e) {
			rs = append(rs, e)
		}
	}

	return rs
}

// Map slice elements
func MapSlice[S, T any](s []T, convert func(t T) S) []S {
	rs := []S{}

	for _, e := range s {
		rs = append(rs, convert(e))
	}

	return rs
}

func First[T any](s []T) T {
	return s[0]
}

func Last[T any](s []T) T {
	return s[len(s)-1]
}

func ContainsSlice[T comparable](s []T, t T) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}
