package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(sum(2, 2))
	resultado, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
	result, err := multiply(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// Multiply
func multiply(a int, b int) (int, error) {
	result := a * b
	if result < 0 {
		return 0, errors.New("Não pode multiplicar negativo")
	}
	return result, nil
}

// Soma
func sum(a int, b int) int {
	return a + b
}

// Division
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Não pode dividir por 0")
	}
	return a / b, nil
}
