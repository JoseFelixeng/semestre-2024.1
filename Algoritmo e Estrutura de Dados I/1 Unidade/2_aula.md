# Aula 2 - LISTAS 
Usamos listas no cotidiano é de comum conhecimento das listas e suas formas utilizações. 

### Caracteristicas das listas 

* Inserir um elemento no fim da lista.
* Ler um elemento em uma posição específica da lista.
* Inserir um elemento em uma posição específica na lista.
* Modificar um elemento em uma posição específica na lista.
* Contar os elementos de uma lista.
* Remover um elemento no fim da lista.
* Remover um elemento em uma posição específica na lista.

Com as caracteristicas escritas a cima podemos vê que um ```Array``` é suficiente para representar as caracteristicas exibidas a cima. Outra opção seria usar as listas ligadas para ligar um dado ponto a outro atráves de ponteiros e assim combinar as caracteristicas e chegar no objetivo.

### ArrayLists
Abaixo podemos vê uma imagem que representa um arrayList.

![Representação array list](/Algoritmo%20e%20Estrutura%20de%20Dados%20I/img/arraylist-apresentacao.png)

Podemos usar arraylists na solução de varios problema porem devemos ter em mente que como qualquer outro metodo, tem seus pontos positivos e negativos, esses que são dados mediante as operações feitas, que podem ser rápidas ou lentas dentro do processo de execução.

Vamos criar um exemplo e organizar os pensamentos para usar o mesmo da melhor forma possível.

```C
struct arraylist {
    int* vetor;
    int qtdade;     
    int capacidade; //capacidade atual do vetor
};

struct arraylist* inicializar(int capacidade) {
    //precisamos alocar memória para o struct
    struct arraylist* lista = (struct arraylist*)malloc(sizeof(struct arraylist));
    //também precisamos alocar memória para o int * (calloc inicializa as posições com 0)
    lista->vetor = (int*)calloc(capacidade, sizeof(int));
    //eh importante inicializar os valores do struct, 
    //pois de outro modo eles iniciarão com lixo (valores utilizados por processos antigos)
    lista->qtdade = 0;
    lista->capacidade = capacidade;
    return lista;
}

```
A forma como esta feito o exemplo acima antes de inialização o arraylist devemos utilizar uma variavel que terá atribuida para si o valor de memoria necessario para criação da mesma, que pode ser uma quantidade muito grande e dessa forma ficar espaços vazios ou pode ser uma quantidade pequena e nessa forma ser necessario novamente o processamento dessa variavel para aumentar o tamanho do array.

#### Inserção no fim da lista

A peculiaridade da inserção é que arrays instanciados não mudam de tamanho naturalmente.  Portanto, sempre que o array estiver totalmente preenchido, uma nova inserção irá precisar criar um novo array com maior capacidade, ou uma realocação que aumente a capacidade do array. Obviamente isto requer que os elementos previamente inseridos no array original sejam copiados para este novo array.


![Representação array list](/Algoritmo%20e%20Estrutura%20de%20Dados%20I/img/arraylist-insercao-fim.gif)


```C

void duplicarCapacidade(struct arraylist* lista){
    int novaLista[2 * lista->capacidade];
    for(int i = 0; i < lista->capacidade; i++){
        novaLista[i] = lista->vetor[i];
    }
    free(lista->vetor);
    lista->vetor = novaLista;
}

void inserirElementoNoFim(struct arraylist* lista, int valor) {
    if (lista->qtdade == lista->capacidade) {
        duplicarCapacidade(lista);
    }
    lista->vetor[lista->qtdade] = valor;
    lista->qtdade++;
}


```

A linguagem C nos provê ferramentas para um gerenciamento rigoroso de memória RAM. A forma mais trivial de aumentar o tamanho de um array é alocando um novo array com maior capacidade e copiando os elementos do array antigo para o novo array, como mostrado acima. Agora vamos usar a função realloc, que é capaz de aumentar o tamanho do bloco de memória alocado, movendo-o para outra posição na memória quando for necessário.

```C

void duplicarCapacidade(struct arraylist* lista) {
    lista->vetor = (int*)realloc(lista->vetor, 2 * lista->capacidade * sizeof(int));
    lista->capacidade = 2 * lista->capacidade;
}

```


Os arrays naturalmente nos provêem a opção de obter um elemento em uma posição específica usando o índice. O único cuidado que devemos ter é verificar se aquele índice é válido.Note também que apesar da lista ter capacidade específica, ela pode ter uma quantidade de elementos menor do que esta capacidade. Na ilustração a seguir, a capacidade usada como exemplo é 10. Então, precisamos ter uma variável para controlar quantos elementos já foram inseridos na lista, de forma a usar este valor para saber se existe algum elemento em uma posição específica. A seguir são apresentadas uma animação que ilustra um ArrayList sendo inicializado, além da obtenção do elemento no índice 2 do ArrayList, e a implementação da função obterElemento.

![Recriando a forma de alocar](/Algoritmo%20e%20Estrutura%20de%20Dados%20I/img/arraylist-obtencao.gif)

```C
int obterElementoEmPosicao(struct arraylist* lista, int posicao) {
    if (posicao >= 0 && posicao < lista->qtdade) {
        return lista->vetor[posicao];
    }
}
```

Para inserir um elemento em uma posição específica precisamos tomar o cuidado de deslocar todos os elementos uma posição à direita, para abrir espaço para o novo elemento. Nesta operação, também precisamos tomar o cuidado de verificar se há espaço no array para inserir este elemento, caso contrário, devemos utilizar a função que duplica a capacidade do array.

Lembre-se que na interface inserirElemento(int valor, int posicao), posição refere-se ao índice da lista. Ou seja, para adicionar na primeira posição, posicao = 0. Note também que precisamos verificar se posição é um valor válido.

A seguir é apresentada uma animação que ilustra um ArrayList sendo preenchido, e no final acontece uma inserção na posição 0.

Obs: o código-fonte da função inserirElementoEmPosicao apresentado na animação está incorreto. O código-fonte apresentado depois da animação está correto.

![alocando novos espaços na memoria e verificando](/Algoritmo%20e%20Estrutura%20de%20Dados%20I/img/arraylist-insercao-posicao.gif)

```C
void inserirElementoEmPosicao(struct arraylist* lista, int valor, int posicao) {
    if (posicao >= 0 && posicao <= lista->qtdade) {
        if (lista->qtdade == lista->capacidade) {
            duplicarCapacidade(lista);
        }

        for (int i = lista->qtdade; i > posicao; i--) {
            lista->vetor[i] = lista->vetor[i - 1];
        }
        lista->vetor[posicao] = valor;
        lista->qtdade++;
    }
}
``` 

Já para atualizamos em um local especifico é necessario que saibamos se o local de alocação é válido.

```C


void atualizarElemento(struct arraylist* lista, int valor, int posicao) {
    if (posicao >= 0 && posicao < lista->qtdade) {
        lista->vetor[posicao] = valor;
    }
}

```

Fazer a contagem da lista é necessario um elemento iterado 

```C
int getTamanho(){
    return tamanho;
}
```

Já para remover um elemento do final da lista é necessario apenas remover o ultimo item.

```C
void removerElemento(){
    tamanho--;
}
```

Para remover um elemento de uma posição específica da lista, precisamos deslocar uma posição à esquerda todos os elementos que estão à direita dessa posição. Com isto, o elemento que antes ocupava esta posição é virtualmente removido.

![A imagem representa  a remoção de um item em uma posição especifica](/Algoritmo%20e%20Estrutura%20de%20Dados%20I/img/arraylist-remocao-posicao.gif)


```C
void removerElementoEmPosicao(struct arraylist* lista, int posicao) {
    if (posicao >= 0 && posicao < lista->qtdade) {
        while (posicao < lista->qtdade - 1) {
            lista->vetor[posicao] = lista->vetor[posicao + 1];
            posicao++;
        }
        lista->qtdade--;
    }
}
```

### Noções de Desempenho
É muito importante estudar os algoritmos e EDs sob a perspectiva de seus desempenhos. Embora neste ponto do curso nós ainda não aprendemos sobre análise assintótica de algoritmos, que é o assunto que trata dos formalismos matemáticos para analisar o desempenho de algoritmos, nós podemos fomentar uma discussão que nos permita ter noções de por quê as funções de inserção, remoção, obtenção e atualização desempenham bem ou mal.

#### Arrays são alocados em blocos contíguos de RAM!
Antes de analisar as funções é preciso destacar a principal propriedade de um array: em sua criação, blocos contíguos de RAM são alocados para ele. A grande vantagem disto é que se tivermos um array de int com 10^6 posições, então é possível acessar o último elemento do array, i.e., o elemento da posição (10^6)-1, de forma instantânea, sem precisar escanear todo o array.

Note que em nossas animações dos algoritmos eu escolhi aleatoriamente que o array estava sempre sendo instanciado no endereço 8239. Como o array é de inteiros, e sabemos que cada int ocupa 4 bytes, então sabemos que podemos obter o elemento da posição (10^6)-1 da seguinte forma: 8239 + posição * 4, que neste caso seria 8239 + (10^6-1) * 4. Obviamente esta lógica funciona para arrays de quaisquer tamanhos.

#### Análise de desempenho das funções
Aqui vamos classisficar o desempenho em instantâneo sempre que não precisarmos percorrer o array para executar uma função, e proporcional ao tamanho do array sempre que a função necessitar escanear todo o array (pelo menos na situação mais extrema, que chamamos de pior caso).

* inserirElementoNoFim: depende da necessidade ou não de duplicação do array; quando não for necessário duplicar o array, o desempenho será instantâneo
* duplicarCapacidade (novo array): desempenho será proporcional ao tamanho do array dado que sera preciso iterar no array antigo para copiar os elementos para o novo array
* duplicarCapacidade (realloc): será instantâneo se a realocação de memória não impuser mudança para um novo endereço, mas também pode ser proporcional ao tamanho do array se for preciso copiar o conteúdo do array anterior para um novo endereço
* obterElemento: instantâneo pois como o bloco de memória alocado para o array é contíguo, é possível calcular o exato endereço de memória do elemento que está sendo buscado
* inserirElementoEmPosicao: pode ser proporcional ao tamanho do array se o elemento for inserido na posição 0, o que causa um deslocamento de todos os elementos do array
* atualizarElemento: instantâneo pelas mesmas razões de obterElemento (é possível calcular o exato endereço de memória do elemento que se deseja atualizar)
* getTamanho: instantâneo
* removerElemento: instantâneo
* removerElementoEmPosicao: proporcional ao tamanho do array pelas mesmas razões de inserirElementoEmPosicao (uma remoção do elemento na posição 0 causa um deslocamento de todos os elementos do array)


### Implementando 

```GO

package list

import "errors"  // Importando por herança função ERRORS

// Criando uma estrutura do tipo ArrayList
type ArrayList struct {
	values   []int //Valores para ser usada para instancias os valores das posições 
	inserted int   // variavel usada para inserir novos valores ao arraylist
}

// Função para criar a arraylist
func (list *ArrayList) Init(size int) error {
    // Testa se o tamanho da arraylist é zero
	if size > 0 {
        // adiciona valor no arraylist 
		list.values = make([]int, size)
		return nil // se tudo der certo o retorno será NULL
	} else {
        // Caso não seja possivel criar um arraylist retorna um Erro.
		return errors.New("Can't init ArrayList with size <= 0")
	}
}

/**
 * Duplica o vetor.
 */
func (list *ArrayList) doubleArray() {
	curSize := len(list.values) // criar uma instancia com o tamanho antigo do arraylist
	doubledValues := make([]int, 2*curSize) // Duplicar o tamanho do array list com a entrada do curSize

	for i := 0; i < curSize; i++ {
        // Atualizando o array list com os valores antingo
		doubledValues[i] = list.values[i]
	}
    // atualiza a lista apos ser duplicada
	list.values = doubledValues
}

/**
 * Adiciona sempre da esquerda para a direita,
 * após o último elemento inserido anteriormente.
 *
 * Melhor caso: Ômega(1)
 *   Just: não precisa duplicar vetor, nem varrer o vetor
 *         do início para encontrar o endereço a ser Add
 * Pior caso: O(n)
 *   Just: duplicar o vetor requer copiar os elementos
 *         para o novo vetor, causando um custo computa-
 *         cional proporcional ao tamanho do vetor (n)
 */
func (list *ArrayList) Add(val int) {
	//verificar se array está lotado
    // Testa se a lista esta lotada, caso seja verdadeiro chama a função para dobrar a lista
	if list.inserted >= len(list.values) {
		list.doubleArray()
	}
    // instancia as listas e inserir os valores 
	list.values[list.inserted] = val
	list.inserted++
}

/**
 * Adiciona elemento na posição especificada, deslocando
 * os elementos à direita dessa posição.
 *
 * Melhor caso: Ômega(1)
 *   Just: não precisa duplicar vetor, nem varrer o vetor
 *         do início para encontrar o endereço a ser Add
 * Pior caso: O(n)
 *   Just: 1) duplicar o vetor requer copiar os elementos
 *         para o novo vetor, causando um custo computa-
 *         cional proporcional ao tamanho do vetor (n)
 *         2) adicionar no início requer deslocar os n
 *            elementos do vetor para a direita
 */
func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.inserted {
		if list.inserted == len(list.values) {
			list.doubleArray()
		}
		for i := list.inserted; i > index; i-- {
			list.values[i] = list.values[i-1]
		}
		list.values[index] = value
		list.inserted++
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't add in arraylist on negative index")
		} else {
			return errors.New("Can't add in arraylist on index out of upper bounds")
		}
	}
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < list.inserted {
		for i := index; i < list.inserted-1; i++ {
			list.values[i] = list.values[i+1]
		}
		list.inserted--
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't remove from arraylist on negative index")
		} else {
			return errors.New("Can't remove from arraylist on index out of upper bounds")
		}
	}
}

func (list *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		return list.values[index], nil
	} else {
		if index < 0 {
			return index, errors.New("Can't get value from arraylist on negative index")
		} else {
			return index, errors.New("Can't get value from arraylist on index out of upper bounds")
		}
	}
}

func (list *ArrayList) Set(value, index int) error {
	if index >= 0 && index < list.inserted {
		list.values[index] = value
		return nil
	} else {
		if index < 0 {
			return errors.New("Can't set value on arraylist on index < 0")
		} else {
			return errors.New("Can't set value on arraylist on index >= arraylist.size")
		}
	}
}

func (list *ArrayList) Size() int {
	return list.inserted
}


```