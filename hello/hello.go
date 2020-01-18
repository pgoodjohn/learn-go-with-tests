package main

import "fmt"

const englishHelloPrefix =  "Hello, "

// Hello is our domain, we want to separate it from the outside world
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}
