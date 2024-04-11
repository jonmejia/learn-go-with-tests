package main

import "fmt"

func Sum(array []int) int{
	var result int = 0
	for _, number := range array{
		result += number
	}
	fmt.Println(result)
	return result
}

func main(){
	Sum([]int{1,2,3})
}
