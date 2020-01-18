package main

import "fmt"

// Languages
const spanish = "Spanish"
const french = "French"

// Prefixes for Languages
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello is our domain, we want to separate it from the outside world
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix 
	}
	return prefix + name

}

func main() {
	fmt.Println(Hello("World", "English"))
}
