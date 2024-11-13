package main

import (
	"fizzbuzz"
	"fmt"
)

func main() {
	for i := range 20 {
		if i % 3 == 0 || i % 5 == 0 {
			fizzbuzzer, err := fizzbuzz.FizzBuzz(i)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%s\n", fizzbuzzer)
			}
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
