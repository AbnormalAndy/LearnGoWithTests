package main

import (
	"fmt"
	"hello_world"
)

func main() {
	var helloworld string
	var helloenglish string
	var hellospanish string
	var hellofrench string


	helloworld = hello_world.Hello("", "")
	helloenglish = hello_world.Hello("awesome apple", "")
	hellospanish = hello_world.Hello("brave banana", "spanish")
	hellofrench = hello_world.Hello("curious cake", "french")


	fmt.Println(helloworld)
	fmt.Println(helloenglish)
	fmt.Println(hellospanish)
	fmt.Println(hellofrench)
}
