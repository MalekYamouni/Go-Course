package main

import "log"

func main() {
	a := 5
	b := 5
	AddNumbers(a, b)
}

func AddNumbers(a, b int) int {
	x := a + b

	log.Println(x)
	return x
}
