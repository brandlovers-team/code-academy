package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	fmt.Println("Soma:", soma(1, 1), ".")
}

func soma(a int, b int) int {
	return a + b
}
