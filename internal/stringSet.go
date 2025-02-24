package internal

// Structure to simulate a String set.
type StringSet struct {
	setMap map[string]struct{}
}

// NewStringSet initializes an empty String set.
func NewStringSet() *StringSet {
	return &StringSet{
		setMap: make(map[string]struct{}),
	}
}

// NewStringSetFromValues returns a set that contains given values.
func NewStringSetFromValues(values []string) *StringSet {
	tempSet := make(map[string]struct{})

	for _, v := range values {
		tempSet[v] = struct{}{}
	}

	return &StringSet{
		setMap: tempSet,
	}
}

// Add adds a string to the set. Returns true if the set did not contain the
// given string, false otherwise.
func (s *StringSet) Add(value string) bool {
	if _, contains := s.setMap[value]; !contains {
		s.setMap[value] = struct{}{}
		return true
	}

	return false
}

// Contains returns true if the set contains the given value, false otherwise.
func (s *StringSet) Contains(value string) bool {
	_, contains := s.setMap[value]
	return contains
}

// Remove removes the given element from the set. Return true if the element was
// contained in the set, false otherwise.
func (s *StringSet) Remove(value string) bool {
	if _, contains := s.setMap[value]; contains {
		delete(s.setMap, value)
		return true
	}

	return false
}

// Length returns the amount of elements contained in the set (its cardinality).
func (s *StringSet) Length() int {
	return len(s.setMap)
}
