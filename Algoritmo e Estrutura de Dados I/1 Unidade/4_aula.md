# Aula 4

### LinkedList

Otimizar a utilização de memoria e melhorar a estabiliade de espaço usado é uma forma de melhorar o desempenho de um software, assim sendo evitar gastos de memoria desnecessarios e espaços vazios é excessial.

<img src="./img/linkedlist-gerenciamento-memoria.gif">

Uma LinkedLists(ou listas ligadas, em pt-br) são listas implementaadas com nós, e cada nó possui um espaço de memória para armazenar o elemento e outro espaço de memoria para armazenar o ponteiro(endereço de memória) para o nó seguinte.

<img src="./img/linkedlist-apresentacao.png">

Uma outra forma de representar o armazenado de dados na forma de um vetor.

<img src="./img/linkedlist-apresentacao-2.png">

Um ponto positivo para as listas ligadas é o fato de que elas não possuem espaços ociosos. Porém para cada espaço novo alocado é necessario mais 4 bytes para alocar o ponteiro que aponta para o proximo nó. Com isso podemos concluir que dependendo da situação a LinkedList pode ser mais eficiente em temos de alocamento de memoria do que um ArrayList por consumir bem menos espaço.

1. Estrututa:
    1.1 Nós
    1.2 Ponteiros
-- A primeira necessaria a ser criada é a estrutura para um nó.

```Go
struct no{
    int val;
    struct no *prox;
};

```