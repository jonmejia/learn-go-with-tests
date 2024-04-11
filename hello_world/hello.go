package main

import (
	"fmt"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string

	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	case "English":
		prefix = "Hello, "
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
