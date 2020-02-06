package main

// House ...
type House struct {
	Name  string
	Price int64
}

// Index ...
func Index(slice []House, predicate func(House) bool) int {
	for i, elem := range slice {
		if predicate(elem) {
			return i
		}
	}
	return -1
}

// All ...
func All(slice []House, predicate func(House) bool) bool {
	return Index(slice, func(h House) bool {
		return !predicate(h)
	}) == -1
}

// Any ...
func Any(slice []House, predicate func(House) bool) bool {
	return Index(slice, predicate) != -1
}

// None ...
func None(slice []House, predicate func(House) bool) bool {
	return Index(slice, predicate) == -1
}

// Finds ...
func Finds(slice []House, predicate func(House) bool) House {
	result := Index(slice, predicate)
	if result == -1 {
		return House{}
	}
	return slice[result]
}

func main() {}
