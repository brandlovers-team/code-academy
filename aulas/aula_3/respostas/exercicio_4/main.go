package main

import (
	"errors"
	"fmt"
)

func main() {
	if res, err := divide(10, 0); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", res)
	}

	if res, err := divide(10, 2); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", res)
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}
