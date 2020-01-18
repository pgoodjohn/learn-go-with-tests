package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello is our domain, we want to separate it from the outside world
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	} 
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
