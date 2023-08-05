package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix

	case "French":
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	if language == "French" {
		return "Bonjour, " + name
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
