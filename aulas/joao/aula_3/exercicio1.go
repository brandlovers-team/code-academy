package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := multiplicar(-10, -5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

}

func divisao(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("não retornar 0")

	}
	return a / b, nil
}
func multiplicar (a float64, b float64) (float64, error) {
	resultado := a*bif resultado < 0{
		return 0, errors.New ("resultado negativo")
		
	}
	return resultado, nil