package main

import "fmt"

const englishHelloPrefix string = "Hello, "
func Hello(name string) string {
	//single quotes are for 'bytes' and 'runes' in golang, cannot be used for string
	return englishHelloPrefix+name
}

func main() {
	fmt.Println(Hello("world"))
}
