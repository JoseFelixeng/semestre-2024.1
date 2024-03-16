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
