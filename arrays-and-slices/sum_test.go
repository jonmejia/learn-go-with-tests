package main
import "testing"

func TestArray(t *testing.T) {
	numbers := []int{5, 5, 5}
	var result int = Sum(numbers)
	var expectedResult int = 15

	if result != expectedResult{
	}
}
