package main

import "fmt"

const englistPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const defaultName = "World"
const spanishLang = "Spanish"
const frenchLang = "French"

// "Hello, world" string
func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	if language == spanishLang {
		return spanishPrefix + name
	}
	if language == frenchLang {
		return frenchPrefix + name
	}
	return englistPrefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
