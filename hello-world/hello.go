package main

import "fmt"

// Defining constants can help improve the performance of the code
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	// Hello() returns a string
	// Types are added after name, e.g. name string

	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Darshan"))
}
