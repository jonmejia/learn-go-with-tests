package main

import "fmt"

const englishHelloPrefix string = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name

	}
	if language == "French" {
		return "Bonjour, " + name

	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
