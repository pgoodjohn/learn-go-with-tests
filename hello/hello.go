package main

import "fmt"

// Languages
const spanish = "Spanish"
const french = "French"

// Prefixes for Languages
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func getPrefixForLanguage(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix 
	default:
		prefix = englishHelloPrefix
	}

	return 
}

// Hello is our domain, we want to separate it from the outside world
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefixForLanguage(language) + name

}

func main() {
	fmt.Println(Hello("World", "English"))
}
