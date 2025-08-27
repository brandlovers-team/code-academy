# Aula 2

# Linha de Comando

## Por que Usar a Linha de Comando?

A linha de comando (também conhecida como terminal, console ou CLI - Command Line Interface) é uma ferramenta poderosa que permite interagir com o computador através de comandos de texto. Dominar a linha de comando é essencial para programadores porque:

- Oferece mais controle sobre o sistema
- Permite automatizar tarefas repetitivas
- É mais rápida que interfaces gráficas para muitas operações
- É necessária para usar muitas ferramentas de desenvolvimento

## Comandos Básicos

| **Comando** | **Descrição** | **Exemplo** |
| --- | --- | --- |
| cd | Mudar de diretório | `cd pasta`ou`cd ..`(voltar) |
| ls (Mac/Linux)dir (Windows) | Listar arquivos e pastas | `ls`ou`ls -la`(detalhes) |
| mkdir | Criar pasta | `mkdir nova_pasta` |
| rm (Mac/Linux) del (Windows) | Remover arquivos | `rm arquivo.txt` |
| rm -rf (Mac/Linux) rmdir (Windows) | Remover pastas | `rm -rf pasta`(cuidado!) |
| touch (Mac/Linux) echo > (Windows) | Criar arquivo vazio | `touch arquivo.txt` |
| open (Mac) start (Windows) | Abrir arquivos | `open arquivo.txt` |
| cat (Mac/Linux) type (Windows) | Mostrar conteúdo de arquivo | `cat arquivo.txt` |
| pwd (Mac/Linux) cd (Windows) | Mostrar diretório atual | `pwd` |

## Dicas para a Linha de Comando

- Use a tecla Tab para autocompletar comandos e nomes de arquivos
- Seta para cima ↑ mostra comandos anteriores
- Ctrl+C cancela o comando atual
- Ctrl+L limpa a tela
- Cuidado com comandos de remoção como `rm -rf` - eles não pedem confirmação!

### Exercícios

1 Crie uma pasta chamada `projetos`, entre nela, e liste seu conteúdo vazio.

- O que fazer: criar uma pasta chamada "projetos", entrar nessa pasta, e verificar que não há arquivos dentro listando o conteúdo.
- Resposta detalhada

2 Dentro de `projetos`, crie uma subpasta `aula2` e um arquivo vazio `anotacoes.txt`.

- O que fazer: dentro da pasta "projetos", criar uma pasta chamada "aula2" e, dentro dela, criar um arquivo chamado "anotacoes.txt". Confirmar que o arquivo aparece ao listar o conteúdo da pasta.
- Resposta detalhada

3 Escreva uma linha de texto no arquivo `anotacoes.txt` e em seguida mostre o conteúdo no terminal.

- O que fazer: adicionar uma frase ao arquivo chamado "anotacoes.txt" que está dentro da pasta "aula2" e, depois, mostrar o conteúdo desse arquivo na tela.
- Resposta detalhada

4 Crie mais dois arquivos `tarefa1.txt` e `tarefa2.txt`, liste todos os arquivos com detalhes e, por fim, abra `tarefa1.txt` pelo sistema.

- O que fazer: dentro da pasta "aula2", criar dois arquivos chamados "tarefa1.txt" e "tarefa2.txt". Em seguida, listar todos os itens da pasta com detalhes e abrir o arquivo "tarefa1.txt" pelo sistema.
- Resposta detalhada

5 Remova o arquivo `tarefa2.txt` e depois remova a pasta `aula2` (somente se ela estiver vazia). Volte ao diretório anterior e mostre o caminho atual.

- O que fazer: apagar o arquivo chamado "tarefa2.txt" dentro de "aula2". Depois, voltar uma pasta (para "projetos"), remover a pasta "aula2" se estiver vazia e, por fim, mostrar qual é a pasta atual.
- Resposta detalhada

# Git e GitHub

## O que é Git?

Git é um sistema de controle de versão que permite rastrear mudanças em arquivos, colaborar com outras pessoas e gerenciar diferentes versões do seu código.

## O que é GitHub?

GitHub é uma plataforma online que hospeda repositórios Git, facilitando o compartilhamento de código, colaboração e gestão de projetos.

## Comandos Básicos do Git

| **Comando** | **Descrição** | **Exemplo** |
| --- | --- | --- |
| git clone | Copia um repositório remoto para sua máquina | `git clone [https://github.com/usuario/repositorio.git](https://github.com/usuario/repositorio.git)` |
| git branch | Lista, cria ou exclui branches | `git branch`(listar)`git branch nova-feature`(criar) |
| git checkout | Muda para outro branch ou commit | `git checkout nome-do-branch` |
| git pull | Baixa atualizações do repositório remoto | `git pull origin main` |
| git push | Envia suas alterações para o repositório remoto | `git push origin main` |
| git add | Adiciona arquivos para serem commitados | `git add .`(todos os arquivos)`git add arquivo.txt`(um arquivo) |
| git commit | Salva suas alterações com uma mensagem | `git commit -m "Adiciona funcionalidade X"` |
| git status | Mostra o estado dos arquivos | `git status` |

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

## Por que Usar Git?

- Mantém um histórico completo de alterações
- Permite trabalhar em paralelo em diferentes features
- Facilita a colaboração em equipe
- Possibilita voltar a versões anteriores do código
- É uma habilidade essencial para programadores

### Atualizar sua chave SSH no GitHub

Se a conexão SSH com o GitHub falhar ou você trocou de máquina, atualize sua chave SSH.

- Verifique se já existe uma chave na sua máquina
- Gere uma nova chave SSH, se necessário
- Adicione a chave ao ssh-agent
- Adicione a chave pública à sua conta no GitHub
- Teste a conexão com: `ssh -T [git@github.com](mailto:git@github.com)`

Guia oficial do GitHub com passo a passo completo:

- Conectar ao GitHub com SSH: [https://docs.github.com/en/authentication/connecting-to-github-with-ssh](https://docs.github.com/en/authentication/connecting-to-github-with-ssh)
- Adicionar uma nova chave SSH à sua conta: [https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)

### Exercícios (Git e GitHub)

1 Clonar o repositório existente e configurar o remoto

- O que fazer: usar o repositório existente "brandlovers-team/code-academy" (SSH) para criar uma cópia local e confirmar que o remoto "origin" aponta para ele.
- Resposta detalhada

2 Criar uma branch de trabalho seguindo o fluxo (feature branch)

- O que fazer: garantir que a branch principal esteja atualizada, criar uma nova branch com o seu nome no padrão feature/aula2-seu-nome e confirmar que você está nela.
- Resposta detalhada

3 Criar a estrutura de pastas e arquivos da aula dentro do repositório

- O que fazer: dentro da pasta raiz do repositório, criar a estrutura de diretórios students/seu-nome/lesson 1/ e, dentro dela, adicionar arquivos de conteúdo da aula como anotacoes.txt, tarefa1.txt e tarefa2.txt. Depois, voltar para a raiz e conferir o que mudou.
- Resposta detalhada

4 Fazer commits pequenos e descritivos (granulares)

- O que fazer: registrar mudanças em etapas pequenas e claras. Primeiro, comitar a estrutura inicial; depois, adicionar arquivos de exercícios; por fim, ajustes de conteúdo. Cada etapa deve ter uma mensagem de commit descritiva.
- Resposta detalhada

5 Enviar a branch e abrir um Pull Request no GitHub

- O que fazer: enviar a sua branch de trabalho para o remoto "origin" e abrir um Pull Request no repositório no GitHub, com título e descrição claros, para que o instrutor revise e faça o merge.
- Resposta detalhada

# Introdução a Go (Golang)

Go é uma linguagem de programação criada pelo Google. Pense nela como uma ferramenta simples e rápida para construir programas. Ela é rápida como um carro de corrida, mas fácil de dirigir como uma bicicleta.

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

- Curso longo de Go para explorar quando quiser: [YouTube – Introdução a Go](https://www.youtube.com/watch?v=YS4e4q9oBaU)
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
	pi := 3.14 // float64

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
		return 0, [errors.New](http://errors.New)("denominador nao pode ser 0")
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
	somar1(&v) // passa o endereço de v
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

- O que fazer: criar uma pasta chamada "go-aula", dentro dela criar um arquivo "main.go" que imprime a mensagem "Hello, World!" na tela, executar o programa e, opcionalmente, gerar um binário chamado "app".
- Resposta detalhada

2) Função de soma

- O que fazer: criar uma função chamada "soma" que receba dois números inteiros e retorne o resultado da soma. No programa principal, chamar essa função e exibir o resultado.
- Resposta detalhada

3) Função de divisão

- O que fazer: criar uma função chamada "divide" que receba dois números (pode usar números com vírgula) e retorne o resultado da divisão. No programa principal, chamar essa função com valores de exemplo e exibir o resultado.
- Resposta detalhada

4) Tratando erro na divisão por zero

- O que fazer: criar uma versão da função "divide" que devolva também um erro quando o denominador for 0. No programa principal, chamar a função, verificar se veio erro e exibir a mensagem de erro se acontecer.
- Resposta detalhada

5) Ler input do terminal e usar nas funções de soma e multiplicação

- O que fazer: no programa principal, ler dois números digitados pela pessoa no terminal. Criar e usar funções chamadas "soma" e "mult" para calcular os resultados e mostrar na tela os valores de soma e multiplicação.
- Resposta detalhada

6) Desafio: calculadora completa

- Requisitos: soma, subtração, multiplicação, divisão com tratamento de erro em divisão por zero, e leitura correta dos dados de entrada.
- Não precisa de solução aqui. Tente organizar com funções e um laço para repetir até o usuário digitar "sair".

# Preparação para Próxima Aula

## O que vimos hoje:

- Linha de comando: navegação, criação de pastas e arquivos, listagem e abertura de arquivos
- Git e GitHub: fluxo de trabalho com branch, commits pequenos, push e abertura de Pull Request
- Estrutura do repositório: criação de students/seu-nome/lesson 1/ com arquivos da aula
- Go (Golang): instalação e validação com `go version`
- Tipos básicos: int, string, float64, arrays e slices, maps
- Funções e structs (introdução)
- Erros como valores e ponteiros

## Tarefas para a próxima aula:

- [ ]  Concluir os exercícios de Linha de Comando (1–5)
- [ ]  Clonar o repo, criar a branch feature/aula2-seu-nome, organizar a pasta students/seu-nome/lesson 1 e abrir um PR
- [ ]  Instalar o Go e validar com `go version`
- [ ]  Resolver os exercícios de Go 1–5 no seu diretório de aula

## Na próxima aula:

- Aprofundar criação e uso de structs
- Trabalhar mais com maps
- Introdução a testes em Go (testes de funções com testing)
