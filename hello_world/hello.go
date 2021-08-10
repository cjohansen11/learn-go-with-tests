package main

import "fmt"

const englishPrefix = "Hello "

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name + "!"
}

func main() {
	fmt.Println(hello("World"))
}
