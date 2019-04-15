package set

import (
	"sort"
	"strconv"
	"strings"
)

// Set is an integer set
type Set struct {
	m map[int]struct{}
}

// NewSet returns a pointer to an empty set
func NewSet() Set {
	m := make(map[int]struct{})
	return Set{m}
}

// NewSetInit returns a pointer to a set with the passed integer
func NewSetInit(i int) (ns Set) {
	ns = NewSet()
	ns.add(i)
	return ns
}

// Add creates a copy of the set and adds the integer if new to set
func (s Set) Add(i int) {
	if s.Contains(i) {
		return
	}
	s.add(i)
}

// Union merges the second set onto the first and returns whether the set changed
func (s Set) Union(os Set) bool {
	b := false
	for i := range os.m {
		if s.Contains(i) {
			continue
		}
		s.add(i)
		b = true
	}
	return b
}

// Copy returns a duplicate set
func (s Set) Copy() (ns Set) {
	ns = NewSet()
	for i := range s.m {
		ns.add(i)
	}
	return ns
}

// Range executes the ranging function for each element
func (s Set) Range(f func(int)) {
	for i := range s.m {
		f(i)
	}
}

// Contains returns whether the set contains the integer
func (s Set) Contains(i int) bool {
	_, ok := s.m[i]
	return ok
}

// IsEmpty returns true if the set size is zero
func (s Set) IsEmpty() bool {
	return len(s.m) == 0
}

// Print returns a string representation of the set
func (s Set) Print() string {
	ints := make([]int, 0)
	for i := range s.m {
		ints = append(ints, i)
	}
	sort.Ints(ints)
	var sb strings.Builder
	sb.WriteString("{ ")
	for _, i := range ints {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteRune(' ')
	}
	sb.WriteString("}")
	return sb.String()
}

func (s Set) add(i int) {
	s.m[i] = struct{}{}
}
