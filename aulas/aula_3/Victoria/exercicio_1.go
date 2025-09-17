package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(soma(2, 2))
	resultado, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
	fmt.Println(multiply(10, -1))
}

// Teste
func soma(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não pode dividir por zero")
	}
	return a / b, nil
}

func multiply(a int, b int) (int, error) {
	if result < 0 {
	}
	return a * b
}
