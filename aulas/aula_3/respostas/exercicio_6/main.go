package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Bem-vindo à Calculadora!")

	for {
		fmt.Println("\n=============================")
		fmt.Println("Calculadora - Escolha uma operação:")
		fmt.Println("1. Soma")
		fmt.Println("2. Subtração")
		fmt.Println("3. Multiplicação")
		fmt.Println("4. Divisão")
		fmt.Println("5. Sair")
		fmt.Println("=============================")
		fmt.Print("Digite sua escolha: ")

		opcao := lerOpcao()

		if opcao == "5" {
			fmt.Println("Obrigado por usar a calculadora! Até logo!")
			break
		}

		a, err := lerFloat("Digite o primeiro número: ")
		if err != nil {
			fmt.Println("Entrada inválida!")
			continue
		}

		b, err := lerFloat("Digite o segundo número: ")
		if err != nil {
			fmt.Println("Entrada inválida!")
			continue
		}

		switch opcao {
		case "1":
			resultado := soma(a, b)
			fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", a, b, resultado)
		case "2":
			resultado := subtracao(a, b)
			fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", a, b, resultado)
		case "3":
			resultado := multiplicacao(a, b)
			fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", a, b, resultado)
		case "4":
			resultado, err := divisao(a, b)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", a, b, resultado)
			}
		default:
			fmt.Println("Opção inválida! Por favor, escolha entre 1 e 5.")
		}
	}
}

func lerFloat(prompt string) (float64, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	linha, _ := reader.ReadString('\n')
	linha = strings.TrimSpace(linha)
	return strconv.ParseFloat(linha, 64)
}

func lerOpcao() string {
	reader := bufio.NewReader(os.Stdin)
	linha, _ := reader.ReadString('\n')
	return strings.TrimSpace(linha)
}

func soma(a, b float64) float64 {
	return a + b
}

func subtracao(a, b float64) float64 {
	return a - b
}

func multiplicacao(a, b float64) float64 {
	return a * b
}

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}
