# Exercício 5 - Ler input do terminal e usar nas funções de soma e multiplicação

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

## Como executar

```bash
go run main.go
```

O programa vai pedir para você digitar dois números e então mostrará a soma e multiplicação deles.