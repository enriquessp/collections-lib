package collections

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GroupedBySlice(t *testing.T) {
	ints := []int{1, 2, 2, 3}
	groupedByItself := GroupedBySlice(ints, func(n int) int { return n })

	assert.Equal(t, 3, len(groupedByItself))
	assert.Equal(t, []int{1}, groupedByItself[1])
	assert.Equal(t, []int{2, 2}, groupedByItself[2])
	assert.Equal(t, []int{3}, groupedByItself[3])
}

func TestForeachSlice(t *testing.T) {
	ints := []int{1, 2, 3}
	ForeachSlice(ints, func(i int) { assert.Positive(t, i) })
}

func TestContaisSlice(t *testing.T) {
	ints := []int{1, 2, 3}

	found := ContainsSlice(ints, 2)
	notFound := ContainsSlice(ints, 9)

	assert.True(t, found)
	assert.False(t, notFound)
}

func TestFilterSlice(t *testing.T) {
	ints := []int{1, 2, 3}

	odds := FilterSlice(ints, func(n int) bool { return (n % 2) == 1 })
	evens := FilterSlice(ints, func(n int) bool { return (n % 2) == 0 })

	assert.Equal(t, 1, len(evens))
	assert.Equal(t, 2, len(odds))
}

func TestMapSlice(t *testing.T) {
	ints := []int{1, 2, 3}
	intsStr := MapSlice(ints, strconv.Itoa)

	assert.ElementsMatch(t, intsStr, []string{"1", "2", "3"})
}

func TestMapSliceToSet(t *testing.T) {
	ints := []int{1, 2, 3}
	strSet := MapSliceToSet(ints, strconv.Itoa).ToSlice()
	assert.ElementsMatch(t, strSet, []string{"1", "2", "3"})

	ints = []int{}
	strSet = MapSliceToSet(ints, strconv.Itoa).ToSlice()
	assert.ElementsMatch(t, strSet, []string{})
}

func TestSliceUnion(t *testing.T) {
	intTests := []SliceIntTestCase{
		{[]int{1, 3, 4, 6}, []int{1, 2, 3, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4}, []int{0, 2, 3, 5}, []int{1, 2, 3, 4, 0, 5}},
		{[]int{}, []int{44, 3, 2}, []int{44, 3, 2}},
	}

	for _, tc := range intTests {
		t.Run("Integer Tests", func(t *testing.T) {
			u := SliceUnion(tc.s1, tc.s2)
			assert.ElementsMatch(t, u, tc.expected)
		})
	}

	stringTests := []SliceStringTestCase{
		{[]string{"a", "zzz", "b"}, []string{"aa", "b", "c"}, []string{"a", "zzz", "b", "aa", "c"}},
		{[]string{}, []string{}, []string{}},
		{[]string{"asd", "3ef"}, []string{}, []string{"asd", "3ef"}},
		{[]string{"asd", "3ef"}, []string{"ooo"}, []string{"asd", "3ef", "ooo"}},
	}

	for _, tc := range stringTests {
		t.Run("String Tests", func(t *testing.T) {
			u := SliceUnion(tc.s1, tc.s2)
			assert.ElementsMatch(t, u, tc.expected)
		})
	}
}

func TestSliceIntersection(t *testing.T) {
	intTests := []SliceIntTestCase{
		{[]int{1, 3, 4, 6}, []int{1, 2, 3, 5}, []int{1, 3}},
		{[]int{1, 2, 3, 4}, []int{0, 2, 3, 5}, []int{2, 3}},
		{[]int{}, []int{44, 3, 2}, []int{}},
	}

	for _, tc := range intTests {
		t.Run("Integer Tests", func(t *testing.T) {
			u := SliceIntersection(tc.s1, tc.s2)
			assert.ElementsMatch(t, u, tc.expected)
		})
	}

	stringTests := []SliceStringTestCase{
		{[]string{"a", "zzz", "b"}, []string{"aa", "b", "c"}, []string{"b"}},
		{[]string{}, []string{}, []string{}},
		{[]string{"asd", "3ef"}, []string{}, []string{}},
		{[]string{"asd", "3ef"}, []string{"3ef"}, []string{"3ef"}},
	}

	for _, tc := range stringTests {
		t.Run("String Tests", func(t *testing.T) {
			u := SliceIntersection(tc.s1, tc.s2)
			assert.ElementsMatch(t, u, tc.expected)
		})
	}
}

func TestSliceComplement(t *testing.T) {
	intTests := []SliceIntTestCase{
		{[]int{1, 3, 4, 6}, []int{1, 2, 3, 5}, []int{2, 5}},
		{[]int{1, 2, 3, 4}, []int{0, 2, 3, 5}, []int{0, 5}},
		{[]int{}, []int{44, 3, 2}, []int{44, 3, 2}},
		{[]int{1, 2, 3, 4}, []int{3, 2, 1}, []int{}},
	}

	for _, tc := range intTests {
		t.Run("Integer Tests", func(t *testing.T) {
			u := SliceComplement(tc.s1, tc.s2)
			assert.ElementsMatch(t, u, tc.expected)
		})
	}

	stringTests := []SliceStringTestCase{
		{[]string{"a", "zzz", "b"}, []string{"aa", "b", "c"}, []string{"aa", "c"}},
		{[]string{"asd", "3ef"}, []string{}, []string{}},
		{[]string{"asd", "3ef"}, []string{"ooo"}, []string{"ooo"}},
		{[]string{}, []string{}, []string{}},
	}

	for _, tc := range stringTests {
		t.Run("String Tests", func(t *testing.T) {
			u := SliceComplement(tc.s1, tc.s2)
			assert.ElementsMatch(t, u, tc.expected)
		})
	}
}

func TestSliceDifference(t *testing.T) {

	intTests := []SliceIntTestCase{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4, 5}, []int{1, 3}},
		{[]int{1, 2, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}},
		{[]int{1, 1, 2, 2, 2, 3, 99}, []int{0, 2}, []int{1, 3, 99}},
		{[]int{}, []int{44, 3, 2}, []int{}},
	}

	for _, tc := range intTests {
		t.Run("Integer Tests", func(t *testing.T) {
			u := SliceDifference(tc.s1, tc.s2)
			assert.ElementsMatch(t, u, tc.expected)
		})
	}

	stringTests := []SliceStringTestCase{
		{[]string{"_", "{", "{{", "37"}, []string{"{{", "{"}, []string{"_", "37"}},
		{[]string{}, []string{}, []string{}},
		{[]string{"asd", "3ef"}, []string{}, []string{"asd", "3ef"}},
	}

	for _, tc := range stringTests {
		t.Run("String Tests", func(t *testing.T) {
			u := SliceDifference(tc.s1, tc.s2)
			assert.ElementsMatch(t, u, tc.expected)
		})
	}
}

func TestGetFirstAndLast(t *testing.T) {

	s := []string{"First", "Middle", "Last"}
	first := First(s)
	last := Last(s)

	assert.Equal(t, "First", first)
	assert.Equal(t, "Last", last)
}
