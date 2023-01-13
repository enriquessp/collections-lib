package collections

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SliceIntTestCase struct {
	s1       []int
	s2       []int
	expected []int
}

type SliceStringTestCase struct {
	s1       []string
	s2       []string
	expected []string
}

func TestFilter(t *testing.T) {
	ints := []int{1, 2, 3}
	strSet := MapSliceToSet(ints, strconv.Itoa)

	onesElements := strSet.Filter(func(s string) bool { return "1" == s })
	assert.Equal(t, 1, len(onesElements))

	onesAndThreesElements := strSet.Filter(func(s string) bool { return "1" == s || "3" == s })
	assert.Equal(t, 2, len(onesAndThreesElements))
}

func TestMap(t *testing.T) {
	intSet := NewSet[int](3)
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)

	strSet := Map(intSet, strconv.Itoa)
	assert.Equal(t, 3, len(strSet))
}
