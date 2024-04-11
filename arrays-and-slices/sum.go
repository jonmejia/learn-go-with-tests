package main

import "fmt"

func Sum(array [3]int) int{
	var result int = 0
	for _, number := range array{
		result += number
	}
	fmt.Println(result)
	return result
}

func main(){
	Sum([3]int{1,2,3})
}
