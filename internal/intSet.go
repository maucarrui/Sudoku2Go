package internal

// Structure to simulate an integer set.
type IntSet struct {

	// There is no built-in data structure to represent a set. To get around this,
	// define a set as a struct map where all the elements inside the map will
	// always point to an empty struct (an empty struct takes 0 bytes on memory).
	setMap map[int]struct{}
}

// NewIntSet initializes an empty integer set.
func NewIntSet() *IntSet {
	return &IntSet{
		setMap: make(map[int]struct{}),
	}
}

// NewIntSetFromValues returns a set that contains given values.
func NewIntSetFromValues(values []int) *IntSet {
	tempSet := make(map[int]struct{})

	for _, v := range values {
		tempSet[v] = struct{}{}
	}

	return &IntSet{
		setMap: tempSet,
	}
}

// Add adds an integer to the set. Returns true if the set did not contain the
// given integer, false otherwise.
func (s *IntSet) Add(value int) bool {
	if _, contains := s.setMap[value]; !contains {
		s.setMap[value] = struct{}{}
		return true
	}

	return false
}

// Contains returns true if the set contains the given value, false otherwise.
func (s *IntSet) Contains(value int) bool {
	_, contains := s.setMap[value]
	return contains
}

// Remove removes the given element from the set. Return true if the element was
// contained in the set, false otherwise.
func (s *IntSet) Remove(value int) bool {
	if _, contains := s.setMap[value]; contains {
		delete(s.setMap, value)
		return true
	}

	return false
}

// Length returns the amount of elements contained in the set (its cardinality).
func (s *IntSet) Length() int {
	return len(s.setMap)
}
