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

// Função para adcionar elemento em uma posição predeterminada
func (list *ArrayList) AddOnIndex(value int, index int) error {
	// testa se o index esta dentro do tamanho correto
	if index >= 0 && index <= len(list.values) {
		if list.inserted == len(list.values) {
			// caso seja fora do array é chamada uma função que duplica o tamanho do array
			list.doubleArray()
		}
		// se o valor estiver dentro do tamanho do vetor o mesmo irá abrir um espaço e adicionar na posição mais a esquerda
		for i := list.inserted; i > index; i-- {
			list.values[i] = list.values[i-1]
		}
		// adciona o valor ao vetor
		list.values[index] = value
		list.inserted++
		// retorna nulo para comprovar que a operaçao foi um sucesso
		return nil
		// se  o index for menor que 0
	} else {
		if index < 0 {
			return errors.New("Não foi possivel adicionar a lista Index negativo")
		} else {
			return errors.New("Não pode adicionar em um valor maior que o index da lista")
		}
	}
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	// testa se o index esta dentro do tamanho correto
	if index >= 0 && index < list.inserted {
		// se o valor estiver dentro do tamanho do vetor o mesmo irá abrir um espaço e adicionar na posição mais a esquerda
		for i := index; i < list.inserted; i++ {
			list.values[i] = list.values[i+1]
		}
		// Removendo da lista
		list.inserted--
		// retorna nulo para comprovar que a operaçao foi um sucesso
		return nil
		// se  o index for menor que 0
	} else {
		if index < 0 {
			return errors.New("Não foi possivel adicionar a lista Index negativo")
		} else {
			return errors.New("Não pode adicionar em um valor maior que o index da lista")
		}
	}
}

func (list *ArrayList) Get(index int) (int, error) {
	// testa se o index esta dentro do tamanho correto
	if index >= 0 && index <= list.inserted {
		// Retorna a posição e o objeto necessario e um valor nulo comprovando que não ouve erro
		return list.values[index], nil
	} else {
		if index < 0 {
			return index, errors.New("Não foi possivel adicionar a lista Index negativo")
		} else {
			return index, errors.New("Não pode adicionar em um valor maior que o index da lista")
		}
	}
}

func (list *ArrayList) Set(value, index int) error {
	// testa se o index esta dentro do tamanho correto
	if index >= 0 && index <= list.inserted {
		// Retorna a posição e o objeto necessario e um valor nulo comprovando que não ouve erro
		list.values[index] = value
		return nil
	} else {
		if index < 0 {
			return errors.New("Não foi possivel adicionar a lista Index negativo")
		} else {
			return errors.New("Não pode adicionar em um valor maior que o index da lista >= arraylist.size")
		}
	}
}
