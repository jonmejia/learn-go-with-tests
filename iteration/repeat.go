package iteration

func Repeat(input string, repeatAmount int) string {
	var output string 
	for i := 0; i <repeatAmount; i++ {
		output += input 
	}
	return output
}
