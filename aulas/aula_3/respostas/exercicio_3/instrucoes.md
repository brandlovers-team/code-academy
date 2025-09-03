# Exercício 3 - Função de divisão

**Objetivo:** Criar uma função que divide dois números decimais.

**Passo a passo:**
- Passo 1: Crie a estrutura básica do programa Go
- Passo 2: Crie uma função chamada "divide":
  - Use float64 como tipo dos parâmetros (para números com vírgula)
  - A função recebe dois parâmetros float64
  - Retorna um float64 (o resultado da divisão)
  - Dentro da função, divida o primeiro pelo segundo e retorne
- Passo 3: Na função main:
  - Chame a função divide com dois números (exemplo: 10.0 e 4.0)
  - Mostre o resultado usando fmt.Println
- Passo 4: Teste o programa com diferentes valores

**Dicas:**
- Use float64 para números decimais em Go
- A divisão em Go usa o símbolo /
- Números decimais em Go usam ponto, não vírgula (3.14, não 3,14)

**Atenção:** Por enquanto, não se preocupe com divisão por zero - faremos isso no próximo exercício!

## Como executar

```bash
go run main.go
```