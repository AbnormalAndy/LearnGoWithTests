package main

import (
	"fmt"
	"hello_world"
)

func main() {
	var helloworld	string
	var helloyou	string


	helloworld = hello_world.HelloWorld()
	helloyou = hello_world.HelloYou("curious cake")


	fmt.Println(helloworld)
	fmt.Println(helloyou)
}
