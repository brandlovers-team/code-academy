# Exercício 2 - Função de soma

**Objetivo:** Criar uma função que soma dois números e mostrar o resultado na tela.

**Passo a passo:**
- Passo 1: Crie um novo arquivo "main.go" ou use o anterior
- Passo 2: Comece com a estrutura básica (package main, import fmt)
- Passo 3: Antes da função main, crie uma função chamada "soma":
  - A função deve receber dois parâmetros do tipo int (números inteiros)
  - Dê nomes aos parâmetros (exemplo: a e b)
  - A função deve retornar um int (o resultado)
  - Dentro da função, some os dois parâmetros e retorne o resultado
- Passo 4: Na função main:
  - Chame a função soma passando dois números (exemplo: 2 e 3)
  - Use fmt.Println para mostrar o resultado na tela
  - Você pode mostrar assim: "Soma: " seguido do resultado
- Passo 5: Execute o programa e verifique se mostra o resultado correto

**Dicas:**
- A sintaxe de função em Go é: `func nomeDaFuncao(parametro1 tipo, parametro2 tipo) tipoRetorno { }`
- Para retornar um valor, use a palavra `return`

## Como executar

```bash
go run main.go
```