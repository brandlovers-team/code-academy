Aula 1
O que é Programação?
Programação é o processo de criar instruções para que um computador execute tarefas específicas. Essas instruções são escritas em linguagens de programação, que são formas estruturadas de comunicação com o computador.

Quando programamos, estamos essencialmente:

Definindo problemas que queremos resolver
Criando soluções lógicas para esses problemas
Traduzindo essas soluções para uma linguagem que o computador entenda
Programar não é apenas sobre escrever código, mas sobre pensar logicamente e resolver problemas de forma estruturada.

O que são Algoritmos?
Algoritmos são sequências de instruções ou passos bem definidos para realizar uma tarefa ou resolver um problema. Pense em um algoritmo como uma receita de culinária:

Uma receita contém ingredientes (dados de entrada)
Instruções passo a passo (processamento)
Resultado final (saída)
Na programação, criamos algoritmos para resolver problemas, e esses algoritmos são implementados usando código.

Estruturas de Controle Básicas
Condicional (if/else)
As estruturas condicionais permitem que o programa tome decisões com base em condições:

# Exemplo em pseudocódigo
se (condição) então
faça isso
senão
faça aquilo
fim se
Loops (for, while)
Loops permitem repetir um bloco de código várias vezes:

# Loop for - executa um número específico de vezes
para i de 1 até 5 faça
imprima i
fim para

# Loop while - executa enquanto uma condição for verdadeira
enquanto (condição) faça
faça algo
fim enquanto
Funções
Funções são blocos de código reutilizáveis que executam uma tarefa específica. Elas nos ajudam a organizar o código e evitar repetições.

# Definindo uma função
função soma(a, b)
retorne a + b
fim função

# Chamando a função
resultado = soma(5, 3)
Funções podem:

Receber parâmetros (dados de entrada)
Processar esses dados
Retornar resultados
Pseudocódigo na Prática
O que é Pseudocódigo?
Pseudocódigo é uma forma de escrever algoritmos usando linguagem natural estruturada, sem se preocupar com a sintaxe específica de uma linguagem de programação.

Exemplo: Receita de Bolo Simples com if e while
Receita:

Receita Simples: Bolo de Chocolate de Caneca (3 minutos no microondas)
Ingredientes:

4 colheres de sopa de farinha de trigo
4 colheres de sopa de açúcar
2 colheres de sopa de chocolate em pó
1 ovo
3 colheres de sopa de leite
3 colheres de sopa de óleo

Modo de Preparo:

Pegue uma caneca grande
Coloque a farinha na caneca
Adicione o açúcar
Adicione o chocolate em pó
Misture bem os ingredientes secos
Quebre o ovo na caneca
Adicione o leite
Adicione o óleo
Misture tudo até ficar homogêneo
Leve ao microondas por 3 minutos
Retire com cuidado (está quente!)
Deixe esfriar um pouco (2 minutos)
Coma!
Resposta correta:

INÍCIO
se caneca nao existe entao
caneca < pegar caneca
fimse

		caneca < mistrurar ingredientes(
			4 colheres de farinha, 
			4 colheres de açúcar, 
			2 colheres de chocolate
		)
		
		caneca < mistrurar ingredientes(
			1 ovo, 
			3 colheres de leite, 
			3 colheres de óleo
		)
    
    microondas < caneca
    
    horario fim microondas < agora + 3 minutos
    ligar microondas
    enquanto agora (menor) horario fim microondas entao:
	    continuar
    fim entao
    desligar microondas
    
	  horario fim espera < agora + 2 minutos
    enquanto agora (menor) horario fim espera entao:
	    continuar
    fim entao
    
    Comer bolo
FIM

função mistrurar ingredientes(lista ingredientes)
enquanto ingredientes nao estao homogeneos entao:   
Misturar ingredientes
fim enquanto

    retornar ingredientes
fim função
Tarefa para a Próxima Aula
Crie um pseudocódigo para uma receita de lasanha mais complexa, utilizando as estruturas if, else, for e while. Pense em diferentes condições que podem ocorrer durante o preparo e como tratá-las no seu algoritmo.
Receita: Lasanha Bolonhesa Simples
Ingredientes:
Molho Bolonhesa:

500g de carne moída
1 cebola média picada
2 dentes de alho picados
1 lata de molho de tomate (340g)
2 colheres de sopa de extrato de tomate
1 xícara de água
Sal, pimenta e orégano a gosto
Azeite para refogar

Molho Branco:

3 colheres de sopa de manteiga
3 colheres de sopa de farinha de trigo
500ml de leite
Sal, pimenta e noz-moscada a gosto

Montagem:

1 pacote de massa para lasanha (direto ao forno)
300g de queijo mussarela fatiado
100g de queijo parmesão ralado
Orégano para polvilhar

Modo de Preparo:
Molho Bolonhesa:

Aqueça o azeite em uma panela grande
Refogue a cebola até ficar transparente
Adicione o alho e refogue por 1 minuto
Acrescente a carne moída
Mexa até a carne perder a cor vermelha
Adicione o molho de tomate e o extrato
Adicione a água
Tempere com sal, pimenta e orégano
Cozinhe por 15 minutos em fogo baixo
Reserve

Molho Branco:

Derreta a manteiga em uma panela
Adicione a farinha de trigo
Mexa por 2 minutos em fogo baixo
Adicione o leite aos poucos, sempre mexendo
Continue mexendo até engrossar
Tempere com sal, pimenta e noz-moscada
Desligue o fogo quando estiver cremoso
Reserve

Montagem:

Pré-aqueça o forno a 180°C
Unte um refratário com manteiga
Coloque uma concha de molho bolonhesa no fundo
Faça a primeira camada de massa
Cubra com molho bolonhesa
Adicione uma camada de molho branco
Coloque fatias de mussarela
Repita as camadas até acabar os ingredientes
A última camada deve ser de molho branco
Cubra com mussarela e parmesão ralado
Polvilhe orégano por cima
Cubra com papel alumínio
Leve ao forno por 30 minutos
Retire o papel alumínio
Deixe mais 10 minutos para gratinar
Retire do forno
Espere 10 minutos antes de cortar
Sirva

Dicas:

Se usar massa que precisa cozinhar, ferva antes seguindo as instruções da embalagem
Pode substituir carne moída por frango desfiado ou proteína de soja
Para ficar mais suculenta, adicione um pouco mais de molho
Teste com um garfo se a massa está macia antes de tirar do forno
Jogo para Praticar: Blockly Maze
Para praticar conceitos básicos de programação de forma visual e divertida, vamos utilizar o jogo Blockly Maze:

https://blockly.games/maze

O que é Blockly?
Blockly é uma biblioteca que adiciona um editor de código visual a aplicativos web e móveis. O editor usa blocos que se encaixam para representar conceitos de código, como variáveis, expressões lógicas, loops, e mais.

Benefícios do Blockly Maze:
Aprende a sequenciar instruções lógicas
Pratica estruturas condicionais (if/else)
Utiliza loops para repetir ações
Desenvolve pensamento algorítmico
Fornece feedback visual imediato
Como Jogar:
Arraste e conecte blocos para criar instruções
Use blocos de movimento para navegar pelo labirinto
Utilize blocos de repetição (loops) para ações repetitivas
Implemente condicionais para tomar decisões com base no caminho
Clique em "Run Program" para executar seu código
Desafio:
Complete pelo menos os 5 primeiros níveis do Blockly Maze antes da próxima aula. Experimente resolver os desafios usando o menor número possível de blocos!

Preparação para Próxima Aula
O que vimos hoje:
Conceitos fundamentais de programação e algoritmos
Estruturas de controle: if/else, for, while
Funções e sua importância
Pseudocódigo na prática
Introdução ao Blockly Maze
Comandos básicos da linha de comando
Introdução ao Git e GitHub
Tarefas para a próxima aula:
Completar pelo menos 5 níveis do Blockly Maze
Criar um pseudocódigo para uma receita de bolo complexa
Praticar os comandos básicos da linha de comando
Criar uma conta no GitHub (se ainda não tiver)
Na próxima aula:
Introdução a uma linguagem de programação real
Variáveis e tipos de dados
Prática de estruturas de controle
Criação do primeiro programa
