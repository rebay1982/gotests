package main

import "fmt"

const french = "French"
const spanish = "Spanish"

const enPrefix = "Hello "
const esPrefix = "Hola "
const frPrefix = "Bonjour "

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return determinePrefix(language) + name
}

func determinePrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frPrefix

	case spanish:
		prefix = esPrefix

	default:
		prefix = enPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("", ""))
}
