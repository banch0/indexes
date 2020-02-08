package main

import (
	"errors"
)

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
func Finds(slice []House, predicate func(House) bool) (interface{}, error) {
	result := Index(slice, predicate)
	if result == -1 {
		return nil, errors.New("not found")
	}
	return slice[result], nil
}

func main() {}
