package set

import (
	"strconv"
	"strings"
)

// Set is an integer set
type Set struct {
	m map[int]struct{}
}

// NewSet returns a pointer to an empty set
func NewSet() *Set {
	m := make(map[int]struct{})
	return &Set{m}
}

// NewSetInit returns a pointer to a set with the passed integer
func NewSetInit(i int) (ns *Set) {
	ns = NewSet()
	ns.add(i)
	return ns
}

// Add creates a copy of the set and adds the value
func (s *Set) Add(i int) (ns *Set) {
	ns = s.Copy()
	ns.add(i)
	return ns
}

// Union returns a set union between the two sets
func (s *Set) Union(os *Set) (ns *Set) {
	ns = s.Copy()
	for i := range os.m {
		ns.add(i)
	}
	return ns
}

// Copy returns a duplicate set
func (s *Set) Copy() (ns *Set) {
	ns = NewSet()
	for i := range s.m {
		ns.add(i)
	}
	return ns
}

// Range executes the ranging function for each element
func (s *Set) Range(f func(int)) {
	for i := range s.m {
		f(i)
	}
}

// IsEmpty returns true if the set size is zero
func (s *Set) IsEmpty() bool {
	return len(s.m) == 0
}

// Print returns a string representation of the set
func (s *Set) Print() string {
	var sb strings.Builder
	for i := range s.m {
		sb.WriteString(strconv.Itoa(i))
	}
	return sb.String()
}

func (s *Set) add(i int) {
	s.m[i] = struct{}{}
}
