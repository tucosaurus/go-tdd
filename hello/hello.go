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
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case frenchLang:
		prefix = frenchPrefix
	case spanishLang:
		prefix = spanishPrefix
	default:
		prefix = englistPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
