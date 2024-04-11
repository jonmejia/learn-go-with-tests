package iteration

import "testing"

func TestIteration (t *testing.T){
	var output string = Repeat("a") 
	const expectedOutput = "aaaa"

	if output != expectedOutput {
		t.Errorf("got %s, expected %s", output,expectedOutput)
	}
}
