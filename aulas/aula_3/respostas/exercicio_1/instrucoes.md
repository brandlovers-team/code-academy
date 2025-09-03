# Exercício 1 - Hello World e execução

**Objetivo:** Criar seu primeiro programa em Go que imprime "Hello, World!" na tela.

**Passo a passo:**
- Passo 1: Abra o terminal
- Passo 2: Crie uma pasta chamada "go-aula" usando o comando `mkdir`
- Passo 3: Entre na pasta criada usando o comando `cd`
- Passo 4: Crie um arquivo chamado "main.go" (pode usar VS Code ou outro editor)
- Passo 5: No arquivo, escreva o código básico de Go:
  - Primeira linha: declare o pacote principal (toda aplicação Go precisa disso)
  - Segunda linha: importe o pacote fmt (usado para imprimir coisas na tela)
  - Crie a função main (onde o programa começa a rodar)
  - Dentro da função, use fmt.Println para imprimir "Hello, World!"
- Passo 6: Salve o arquivo
- Passo 7: No terminal, execute o programa com `go run main.go`
- Passo 8 (Opcional): Compile o programa para criar um executável:
  - Use `go build -o app` para criar um arquivo executável chamado "app"
  - Execute o arquivo criado com `./app` (Mac/Linux) ou `.\app.exe` (Windows)

**Dica:** Lembre-se que Go é sensível a maiúsculas e minúsculas!

## Comandos práticos

```bash
# Criar a pasta e entrar nela
mkdir -p go-aula
cd go-aula
```

## Execute e gere o binário

```bash
# Executar o programa
go run main.go

# Gerar o binário
go build -o app

# Executar o binário gerado
./app   # no Windows: .\app.exe
```