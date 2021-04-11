package main

import (
	"fmt"
)

const spanishLanguage = "Spanish"
const frenchLanguage = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Greeting(name string, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name + "!"
}

func greetingPrefix(language string) (prefix string) {

	switch language {

	case frenchLanguage:
		prefix = frenchHelloPrefix
	case spanishLanguage:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Greeting("", ""))
	fmt.Println(Greeting("Paulo", "Spanish"))
	fmt.Println(Greeting("Gerard", "English"))
	fmt.Println(Greeting("Pierre", "French"))
}
