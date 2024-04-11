package integers

import "testing"

func TestAdder (t *testing.T){
	var sum int = Add(2,2)
	var expected int = 4

	if sum != expected{
		t.Errorf("Expected %q, but got %q", expected, sum)
	}
} 
