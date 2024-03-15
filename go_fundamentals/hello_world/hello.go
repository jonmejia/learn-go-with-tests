package main

import "fmt"

func Hello(name string) string {
	//single quotes are for 'bytes' and 'runes' in golang, cannot be used for string
	return "Hello, "+name
}

func main() {
	fmt.Println(Hello("world"))
}
