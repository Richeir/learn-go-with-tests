package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == ""{
		name = "world"
	}

	return greetingPrefix(name, language)
}

func greetingPrefix(name string, language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix + name
	case french:
		prefix = frenchHelloPrefix + name
	default:
		prefix = englishHelloPrefix + name
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
