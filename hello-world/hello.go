package main

import "fmt"

// Defining constants can help improve the performance of the code
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	// Hello() returns a string
	// Types are added after name, e.g. name string

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Private functions start with lower case
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// No need to return variable when it's already defined at the beginning
	return
}

func main() {
	fmt.Println(Hello("Darshan", ""))
}
