package tests

import (
	"testing"

	arraysslices "github.com/mbilaljawwad/golang-channels-learning/internal/arrays_slices"
)

func TestSum(t *testing.T) {
	t.Run("Sum with 2 integers", func(t *testing.T) {
		x, y := 3, 4
		result := arraysslices.Sum(x, y)
		expected := 7

		if result != expected {
			t.Errorf("Result: %d, Expected: %d", result, expected)
		}
	})

	t.Run("Sum with multiple integers", func(t *testing.T) {
		result := arraysslices.Sum(1, 2, 3, 4, 5, 6, 7)
		expected := 28

		if result != expected {
			t.Errorf("Result: %d, Expected: %d", result, expected)
		}
	})

	t.Run("Sum list with size of 5", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		result := arraysslices.SumList(list)
		expected := 15

		if result != expected {
			t.Errorf("Result: %d, Expected: %d", result, expected)
		}
	})
}
