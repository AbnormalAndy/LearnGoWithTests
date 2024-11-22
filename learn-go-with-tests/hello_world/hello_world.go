package hello_world

import (
	"strings"
)

const (
	spanish = "spanish"
	french = "french"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)


func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + strings.Title(name) + "!"
}
	
	
func greetingPrefix(language string) (prefix string) {
	switch strings.ToLower(language) {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}


