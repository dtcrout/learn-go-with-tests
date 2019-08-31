package main

import "fmt"

func Hello(name string) string {
	// Hello() returns a string
	// types are added after name, e.g. name string
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
