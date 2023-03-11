package internal

import (
	"testing"
)

// TestNewIntSet tests if NewIntSet returns an empty integer set.
func TestNewIntSet(t *testing.T) {
	A := NewIntSet()

	// Set A should be empty.
	if A.Length() != 0 {
		t.Errorf(
			"Expected an empty set, but it has length: %d.",
			A.Length())
	}
}

// countUniqueValues is an auxiliary functions which counts the number of unique
// integers contained in an array.
func countUniqueValues(values []int) int {
	intMap := make(map[int]struct{})

	uniqueVals := 0
	for _, val := range values {
		if _, contains := intMap[val]; !contains {
			uniqueVals++
			intMap[val] = struct{}{}
		}
	}

	return uniqueVals
}

// TestNewIntSetFromValues tests if TestNewIntSetFromValues returns an integer
// set which contains the given values.
func TestNewIntSetFromValues(t *testing.T) {
	// Array of values with no repetitions.
	test_values := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 2, 3, 4, 4, 5, 6, 7},
	}

	for _, values := range test_values {
		A := NewIntSetFromValues(values)

		// Check that set A has the correct length.
		if A.Length() != countUniqueValues(values) {
			t.Errorf(
				"Expected set of length %d., but got set of length %d.",
				countUniqueValues(values),
				A.Length())
		}

		// Check that the values were added correctly.
		for _, val := range values {
			if A.Add(val) {
				t.Errorf("Expected set to contain value: %d.", val)
			}
		}
	}
}

// TestAdd tests if Add correctly adds an integer to the int set.
func TestAdd(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	A := NewIntSet()

	length := A.Length()

	for _, val := range values {
		// Add should return true as the value is not contained in the set.
		if !A.Add(val) {
			t.Errorf("Expected set to not contain value: %d.", val)
		}

		// Adding the same element again should return false.
		if A.Add(val) {
			t.Errorf("Expected set to contain value: %d.", val)
		}

		// Increase the expected length of the set.
		length++

		// Check that the length increases accordingly.
		if A.Length() != length {
			t.Errorf(
				"Expected set to have length %d, but it has length: %d.",
				length,
				A.Length())
		}
	}
}

// TestRemove tests if Remove correctly removes an integer of the int set.
func TestRemove(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	A := NewIntSetFromValues(values)

	length := A.Length()

	for _, val := range values {
		if !A.Remove(val) {
			t.Errorf("Expected set to contain value: %d.", val)
		}

		if A.Remove(val) {
			t.Errorf("Expected set to not contain value: %d", val)
		}

		// Decrease the expected length of the set.
		length--

		// Check that the length increases accordingly.
		if A.Length() != length {
			t.Errorf(
				"Expected set to have length %d, but it has length %d.",
				length,
				A.Length())
		}
	}
}
