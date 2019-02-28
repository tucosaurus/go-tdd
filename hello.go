package main

import "fmt"

const englistPrefix = "Hello, "

// "Hello, world" string
func Hello(name string) string {
	if name == "" {
		return englistPrefix + "World"
	}
	return englistPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
