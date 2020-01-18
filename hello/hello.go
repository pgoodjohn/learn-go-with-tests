package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello is our domain, we want to separate it from the outside world
func Hello(name string, language string) string {
	var prefix string
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		prefix = spanishHelloPrefix
	} else {
		prefix = englishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
