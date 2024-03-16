// Construindo um arrayList importanto por herança a classe list

package list

// Importando Biblioteca de Erros
import "errors"

type ArrayList struct {
	values   []int
	inserted int
}

// Criando uma array list
func (list *ArrayList) Init(size int) error {
	//teste para o tamanho da lista
	if size > 0 {
		//Atribuindo o valor a lista
		list.values = make([]int, size)
		// Retorno nulo para comprovar que foi feita a lista
		return nil
	} else {
		// retorna que não foi possível iniciar a lista
		return errors.New("Can't init ArrayList with size <=0")
	}
}

// Duplicar o Vetor Caso seja necessario;

func (list *ArrayList) doubleArray() {
	// Usado para pegar o tamanho da lista
	curSize := len(list.values)
	// Duplicando valores da lista
	doubledValues := make([]int, 2*curSize)
	for i := 0; i < curSize; i++ {
		doubledValues[i] = list.values[i]
	}
	// Adiciona o tamanho duplicado a lista
	list.values = doubledValues
}

/**
1. Adicionar sempre da esquerda para a direita, após o ultimo elemento inserido anteriormente.
2. Melhor Caso: ômega(1)
Just: não precisa duplicar o vetor, nem varrer o vetor do início para o endereço a ser Add
3. Pior caso: O(n)
Just: Duplicar o vetor requer copiar os elementos para o novo vetor, causando um custo computacional
proporcional ao tamanho do vetor (n).

**/

// Função para adicionar elemento
func (list *ArrayList) Add(val int) {
	// Verificar se array está lotado
	if list.inserted >= len(list.values) {
		list.doubleArray()
	}
	// Após verificar, identifica a posição e inseri o novo item
	list.values[list.inserted] = val
	list.inserted++
}

/**
1. Adiciona elemento na posição especificada, deslocando os elementos à direita dessa posição
2. Melhor caso: ômega(1)
just: Não precisa duplicar o vetor, nem varrer o vetor do ínicio para encontrar o endereço a ser ADD.
3. Pior caso: O(n):
Just:
	1) Duplicar o vetor requer copiar os elementos para o novo vetor, causando um custo computacional proporcional
	ao tamanho do vetor(n);
	2) Adicionar no início requer deslocar os n elementos para a direita

**/
