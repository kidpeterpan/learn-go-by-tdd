package arrays_and_slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T)  {

	t.Run("Sum of 1,2,3,4,5 is 15", func(t *testing.T) {
		numbers := [5]int{1,2,3,4,5}
		expected := 15

		actual := SumWithArray(numbers)

		assert.Equal(t,expected,actual,"want %d, got %d, given %v",expected,actual,numbers)
	})

	t.Run("Sum of 1,2,3 is 6", func(t *testing.T) {
		numbers := []int{1,2,3}
		expected := 6

		actual := SumWithSlice(numbers)

		assert.Equal(t,expected,actual,"want %d, got %d, given %v",expected,actual,numbers)
	})

	// use "go test -cover" to check test coverage

	t.Run("Sum all 2 slice should give the correct result", func(t *testing.T) {
		numbers0 := []int{1,2,3}
		numbers1 := []int{4,5,6}
		expected := []int{6,15}

		actual := SumAll(numbers0,numbers1)

		assert.Equal(t,expected, actual,"want %v, got %v")
	})
}



