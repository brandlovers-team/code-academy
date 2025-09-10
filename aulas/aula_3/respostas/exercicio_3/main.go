package main

import "fmt"

func main() {
	fmt.Println("Divisão:", divide(10, 4))
}

func divide(a, b float64) float64 {
	return a / b
}
