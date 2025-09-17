package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := divisao(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
	{
		resultado, err := multiplicacao(10, -5)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resultado)
		}

	}
}

func divisao(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não coloque 0, animal")
	}
	return a / b, nil
}

func soma(a int, b int) int {
	return a + b
}

func multiplicacao(a float64, b float64) (float64, error) {
	resultado := a * b
	if resultado < 0 {
		return 0, errors.New("resultado negativo")
	}
	return resultado, nil
}
