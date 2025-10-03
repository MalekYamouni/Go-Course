package main

import (
	"fmt"
)

func main() {

	text := "ich liebe dich wirklich sehr ich hoffe wir sehen und bald wieder."

	result, err := WordCounter(text)
	if err != nil {
		fmt.Println("Ihr Wort wurde nicht gefunden.")
		return
	}

	fmt.Println("ich kommt vor:", result["ich"])
}

func AddNumbers(a, b int) int {
	x := a + b
	return x
}

func DivideNumbers(a, b int) int {
	x := a / b
	return x
}

func Counter() {
	for i := 0; i <= 10; i++ {

		if i%2 >= 1 {
			continue
		}
		fmt.Println(i)
	}
}

func FizzBuzz() {
	for i := 1; i <= 60; i++ {

		if i%15 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		fmt.Println(i)
	}
}
