package hello_world

import (
	"strings"
)


func HelloWorld() string {
	return "Hello, world!"
}


func HelloYou(name string) string {
	you := "Hello, " + strings.Title(name) + "!"
	return you
}
