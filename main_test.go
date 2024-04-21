package main

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	expected := 15
	actual := sum(numbers)

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
