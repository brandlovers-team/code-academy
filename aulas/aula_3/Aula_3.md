# Aula 3

# Recap Git e GitHub

## O que é Git?

Git é um sistema de controle de versão que permite rastrear mudanças em arquivos, colaborar com outras pessoas e
gerenciar diferentes versões do seu código.

## O que é GitHub?

GitHub é uma plataforma online que hospeda repositórios Git, facilitando o compartilhamento de código, colaboração e
gestão de projetos.

## Comandos Básicos do Git

| **Comando**  | **Descrição**                                   | **Exemplo**                                                                                          |
|--------------|-------------------------------------------------|------------------------------------------------------------------------------------------------------|
| git clone    | Copia um repositório remoto para sua máquina    | `git clone [https://github.com/usuario/repositorio.git](https://github.com/usuario/repositorio.git)` |
| git branch   | Lista, cria ou exclui branches                  | `git branch`(listar)`git branch nova-feature`(criar)                                                 |
| git checkout | Muda para outro branch ou commit                | `git checkout nome-do-branch`                                                                        |
| git pull     | Baixa atualizações do repositório remoto        | `git pull origin main`                                                                               |
| git push     | Envia suas alterações para o repositório remoto | `git push origin main`                                                                               |
| git add      | Adiciona arquivos para serem commitados         | `git add .`(todos os arquivos)`git add arquivo.txt`(um arquivo)                                      |
| git commit   | Salva suas alterações com uma mensagem          | `git commit -m "Adiciona funcionalidade X"`                                                          |
| git status   | Mostra o estado dos arquivos                    | `git status`                                                                                         |

## Fluxo Básico de Trabalho com Git

```bash
# Clonar o repositório
git clone [https://github.com/usuario/repositorio.git](https://github.com/usuario/repositorio.git)

# Entrar na pasta do repositório
cd repositorio

# Criar um novo branch para sua feature
git branch minha-nova-feature

# Mudar para o novo branch
git checkout minha-nova-feature

# [Faça suas alterações nos arquivos]

# Adicionar as alterações
git add .

# Commitar as alterações
git commit -m "Adiciona nova funcionalidade"

# Enviar as alterações para o GitHub
git push origin minha-nova-feature
```

# Introdução a Go (Golang)

Go é uma linguagem de programação criada pelo Google. Pense nela como uma ferramenta simples e rápida para construir
programas. Ela é rápida como um carro de corrida, mas fácil de dirigir como uma bicicleta.

## Instalação do Go (macOS e Windows)

- Passo 1: Baixar o instalador no site oficial: [https://go.dev/dl](https://go.dev/dl)
- Passo 2: Escolher o arquivo certo para seu sistema
    - macOS: arquivo .pkg
    - Windows: arquivo .msi
- Passo 3: Abrir o arquivo baixado e clicar em “Next/Continuar” até finalizar.
- Passo 4: Fechar e abrir novamente o Terminal ou Prompt de Comando.
- Passo 5: Validar a instalação:
    - Rode: `go version`
    - Deve aparecer algo como: `go version go1.x.x darwin/amd64` (mac) ou `windows/amd64` (Windows)
- Se não funcionar:
    - Reinicie o computador e tente `go version` de novo.
    - No Windows, teste no “Prompt de Comando” ou “PowerShell”.

## Vídeo de apoio

- Curso longo de Go para explorar quando
  quiser: [YouTube – Introdução a Go](https://www.youtube.com/watch?v=YS4e4q9oBaU)
    - Não precisa assistir tudo agora. Use como referência quando surgir dúvida.

## Por que usar Go para scripts nativos no macOS e Windows?

- Simples de ler: o código parece uma receita clara.
- Rápida: programas abrem e rodam muito rápido.
- Binário único: você compila e ganha um arquivo só, fácil de enviar e rodar.
- Multiplataforma: compila para macOS, Windows e Linux com poucos ajustes.
- Confiável: checagem de erros explícita deixa tudo mais seguro.

## Tipos básicos (como se fossem caixas para guardar coisas)

- int: números inteiros. Exemplo: 1, 2, 42.
- string: texto entre aspas. Exemplo: "Olá".
- float64: números com vírgula. Exemplo: 3.14.
- array: uma fileira de caixas com tamanho fixo. Ex: [3]int guarda 3 inteiros.
- slice: uma fileira elástica, que cresce. Ex: []int.
- map: um dicionário que liga chave a valor. Ex: map[string]int.

```go
package main

import "fmt"

func main() {
	var idade int = 10
	nome := "Ana" // inferência de tipo
	pi := 3.14    // float64

	// array com tamanho fixo 3
	var a [3]int = [3]int{1, 2, 3}
	// slice (tamanho flexível)
	numeros := []int{4, 5, 6}
	// map de string para int
	pontos := map[string]int{"Ana": 10, "Joao": 8}

	fmt.Println(idade, nome, pi, a, numeros, pontos)
}
```

## Funções: pequenos robôs que fazem tarefas

- Você dá um nome, entradas e, se quiser, um retorno.

```go
func soma(a int, b int) int {
return a + b
}
```

## Structs: caixinhas com vários campos

- Pense em um "cartão de registro" com várias informações juntas.

```go
type Pessoa struct {
Nome string
Idade int
}

func main() {
p := Pessoa{Nome: "Ana", Idade: 10}
fmt.Println(p.Nome, p.Idade)
}
```

## Erros em Go: são valores

- Em Go, erro é um valor especial (do tipo error). Você checa se é nil. Se não for nil, deu erro.

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, [
		errors.New](http: //errors.New)("denominador nao pode ser 0")
}
return a / b, nil
}

func main() {
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", res)
}
```

## Ponteiros: setinhas que apontam para lugares na memória

- Em vez de copiar a coisa, você entrega o endereço dela.
- Útil para mudar valores dentro de funções.

```go
package main

import "fmt"

func somar1(x *int) { // x é um ponteiro para int
	*x = *x + 1 // altera o valor no endereço
}

func main() {
	v := 10
	somar1(&v)     // passa o endereço de v
	fmt.Println(v) // 11
}
```

- Em vez de copiar a coisa, você entrega o endereço dela.
- Útil para mudar valores dentro de funções.

```go
func somar1(x *int) { // x é um ponteiro para int
*x = *x + 1 // altera o valor no endereço
}

func main() {
v := 10
somar1(&v) // passa o endereço de v
fmt.Println(v) // 11
}
```

### Exercícios de Go

1) Hello World e execução

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


2) Função de soma

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


3) Função de divisão

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


4) Tratando erro na divisão por zero

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


5) Ler input do terminal e usar nas funções de soma e multiplicação

**Objetivo:** Criar um programa interativo que pede números ao usuário e mostra soma e multiplicação.

**Passo a passo:**
- Passo 1: Importe os pacotes necessários:
  - "fmt" para imprimir e ler
  - "bufio" para ler linhas do terminal
  - "os" para acessar o terminal
  - "strconv" para converter texto em número
  - "strings" para limpar espaços
- Passo 2: Crie duas funções simples:
  - função "soma": recebe dois int, retorna a soma
  - função "mult": recebe dois int, retorna a multiplicação
- Passo 3: Crie uma função auxiliar "lerInt" que:
  - Recebe uma mensagem para mostrar ao usuário (tipo string)
  - Mostra a mensagem com fmt.Print
  - Cria um reader com bufio.NewReader(os.Stdin)
  - Lê uma linha com reader.ReadString('\n')
  - Remove espaços em branco com strings.TrimSpace
  - Converte o texto para número com strconv.Atoi
  - Retorna o número e um possível erro
- Passo 4: Na função main:
  - Use lerInt para pedir o primeiro número
  - Verifique se houve erro na leitura
  - Use lerInt para pedir o segundo número
  - Verifique se houve erro na leitura
  - Chame as funções soma e mult com os números lidos
  - Mostre os dois resultados

**Dicas:**
- Ao ler do terminal, sempre verifique erros
- Use mensagens claras como "Digite o primeiro número: "
- Se houver erro na leitura, mostre "Entrada inválida" e termine o programa


6) Desafio para casa: calculadora completa

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

# Preparação para Próxima Aula

## O que vimos hoje:

- Git e GitHub: fluxo de trabalho com branch, commits pequenos, push e abertura de Pull Request
- Estrutura do repositório: criação de students/seu-nome/lesson 1/ com arquivos da aula
- Go (Golang): instalação e validação com `go version`
- Tipos básicos: int, string, float64, arrays e slices, maps
- Funções e structs (introdução)
- Erros como valores e ponteiros

## Tarefas para a próxima aula:

- [ ]  Concluir os exercícios de Linha de Comando (1–6)
- [ ]  Clonar o repo, criar a branch feature/aula2-seu-nome, organizar a pasta students/seu-nome/lesson 1 e abrir um PR
- [ ]  Instalar o Go e validar com `go version`
- [ ]  Resolver os exercícios de Go 1–6 no seu diretório de aula

## Na próxima aula:

- Aprofundar criação e uso de structs
- Trabalhar mais com maps
- Introdução a testes em Go (testes de funções com testing)
