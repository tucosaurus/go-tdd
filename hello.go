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

	prefix := englistPrefix

	switch language {
	case frenchLang:
		prefix = frenchPrefix
	case spanishLang:
		prefix = spanishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
