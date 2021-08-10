package main

import "fmt"

const spanish = "spanish"
const french = "french"
const english = "english"
const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bonjour "

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(hello("World", "english"))
}
