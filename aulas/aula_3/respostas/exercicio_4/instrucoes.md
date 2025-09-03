# Exercício 4 - Tratando erro na divisão por zero

**Objetivo:** Melhorar a função de divisão para lidar com erro quando alguém tenta dividir por zero.

**Passo a passo:**
- Passo 1: Comece com a estrutura básica e importe dois pacotes:
  - "fmt" para imprimir
  - "errors" para criar mensagens de erro
- Passo 2: Crie uma função "divide" melhorada:
  - Recebe dois float64 como antes
  - MAS AGORA retorna DOIS valores: (float64, error)
  - Dentro da função:
    - Primeiro, verifique se o segundo número é 0
    - Se for 0, retorne 0 e um erro com mensagem "divisão por zero"
    - Se não for 0, faça a divisão normal e retorne o resultado e nil
- Passo 3: Na função main:
  - Chame a função divide e receba os DOIS valores de retorno
  - Use a sintaxe: resultado, erro := divide(10, 0)
  - Verifique se o erro não é nil (isso significa que houve erro)
  - Se houver erro, mostre a mensagem de erro
  - Se não houver erro, mostre o resultado
- Passo 4: Teste duas vezes:
  - Uma vez dividindo por 0 (deve mostrar erro)
  - Uma vez dividindo normalmente (deve mostrar resultado)

**Dicas:**
- Para criar um erro: `errors.New("sua mensagem")`
- nil em Go significa "nada" ou "vazio"
- Use if err != nil para verificar se houve erro

## Como executar

```bash
go run main.go
```