package main

import "fmt"

func Hello() string {
	//single quotes are for 'bytes' and 'runes' in golang, cannot be used for string
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}
