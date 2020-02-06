package main

// Index ...
func Index(slice []int, predicate func(int) bool) int {
	for i, elem := range slice {
		if predicate(elem) {
			return i
		}
	}
	return -1
}

// All ...
func All(slice []int, predicate func(int) bool) bool {
	return Index(slice, func(i int) bool {
		return !predicate(i)
	}) == -1
}

// Any ...
func Any(slice []int, predicate func(int) bool) bool {
	return Index(slice, predicate) != -1
}

// None ...
func None(slice []int, predicate func(int) bool) bool {
	return Index(slice, predicate) == -1
}

// Find ...
func Find(slice []int, predicate func(int) bool) int {
	result := Index(slice, predicate)
	if result == -1 {
		return 0
	}
	return slice[result]
}

func main() {}

func nulls(i int) bool {
	return true
}
