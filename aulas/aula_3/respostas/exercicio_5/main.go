package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, err := lerInt("Digite o primeiro número: ")
	if err != nil {
		fmt.Println("Entrada inválida")
		return
	}
	b, err := lerInt("Digite o segundo número: ")
	if err != nil {
		fmt.Println("Entrada inválida")
		return
	}

	fmt.Println("Soma:", soma(a, b))
	fmt.Println("Multiplicação:", mult(a, b))
}

func lerInt(prompt string) (int, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	linha, _ := reader.ReadString('\n')
	linha = strings.TrimSpace(linha)
	return strconv.Atoi(linha)
}

func soma(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}
