package main

import "fmt"

const englistPrefix = "Hello, "
const spanishPrefix = "Hola, "
const defaultName = "World"
const spanishLang = "Spanish"

// "Hello, world" string
func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	if language == spanishLang {
		return spanishPrefix + name
	}
	return englistPrefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
