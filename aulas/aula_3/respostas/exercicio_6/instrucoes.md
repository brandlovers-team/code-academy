# Exercício 6 - Desafio para casa: calculadora completa

**Objetivo:** Criar uma calculadora interativa completa que roda até o usuário decidir sair.

**Requisitos do programa:**
- Deve ter 4 operações: soma (+), subtração (-), multiplicação (*) e divisão (/)
- Deve tratar erro de divisão por zero
- Deve permitir fazer várias operações seguidas
- Deve ter opção de sair do programa

**Passo a passo sugerido:**
- Passo 1: Crie funções separadas para cada operação:
  - soma(a, b float64) float64
  - subtracao(a, b float64) float64
  - multiplicacao(a, b float64) float64
  - divisao(a, b float64) (float64, error) - com tratamento de erro
- Passo 2: Crie um loop infinito na função main usando `for { }`
- Passo 3: Dentro do loop, mostre um menu:
  - "Calculadora - Escolha uma operação:"
  - "1. Soma"
  - "2. Subtração"
  - "3. Multiplicação"
  - "4. Divisão"
  - "5. Sair"
- Passo 4: Leia a escolha do usuário
- Passo 5: Se escolheu sair (5), use `break` para sair do loop
- Passo 6: Para qualquer outra opção:
  - Peça o primeiro número
  - Peça o segundo número
  - Execute a operação correspondente
  - Mostre o resultado
  - Se for divisão e houver erro, mostre a mensagem de erro
- Passo 7: Após mostrar o resultado, o loop volta ao início (menu)

**Dicas extras:**
- Use switch/case ou if/else para lidar com as opções do menu
- Limpe a tela entre operações se quiser (opcional)
- Adicione validações extras: verificar se a entrada é realmente um número
- Mensagem de despedida quando o usuário escolher sair

## Como executar

```bash
go run main.go
```

O programa mostrará um menu interativo onde você pode escolher operações matemáticas e realizar cálculos repetidamente até escolher sair.