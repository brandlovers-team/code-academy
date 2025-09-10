package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	fmt.Println(soma(6, 5))
}

func soma(a int, b int) int {
	return a + b
}
