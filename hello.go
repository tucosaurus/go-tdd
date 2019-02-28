package main

import "fmt"

const englistPrefix = "Hello, "

// "Hello, world" string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}
	return englistPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
