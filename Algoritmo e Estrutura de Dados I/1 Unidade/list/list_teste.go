// Usado para testar as implementações de uma lista

package list

import "testing"

// usando como função de teste

var size int // variavel usada para medir o tamanho

// var lists [3]Ilist
var lists [1]Ilist

// função usada para criar uma lista
// := usado para criar e instancia uma nova variavel

func createLists(size int) {
	arraylist := &ArrayList{}   // usado para criar novos arraylist
	(*arraylist).Init(size)     // Cria um ponteiro chamado arraylist e inicializa o mesmo com o tamanho size
	lists = [1]Ilist{arraylist} // A variavel lista esta recebendo um ponteiro de listas
	// lists = [3]Ilist{arraylist, &linkedList{}, &DoublyLinkedList{}}
}

// função usada para deleta uma lista
// nil no go representa ao nulo.

func deleteLists() {
	// nesse caso, createLists sozinho resolve o problema]
	// no entanto, deixo o template aqui para ser usado em outros testes
	lists[0] = nil
	// lists[1] = nil
	// lists[2] = nil
}

//função usada para a configuração de listas
func setupTest() func() {
	// antes de cada teste
	size = 10         // atribuido a variavel size o valor 10
	createLists(size) // cria uma lista de tamanho 10

	// depois de cada teste
	return func() {
		deleteLists() // usada para retornar a função delete list
	}
}

// função usada para inicias os teste de adcionar itens
//  o identificador _ é usado para "descartar" valores retornados por funções ou atribuídos a variáveis, evitando assim que o compilador gere erros por variáveis não utilizadas.

func TestAdd(t *testing.T) {
	defer setupTest()()          // usado para chamar a função setupTest
	for _, list := range lists { // usado para criar uma variavel lista e ao mesmo tempo atribuir um valor na memoria para a mesma
		for i := 0; i < size; i++ { // usado para criar e instacia a variavel i que irá ser atualizada baseado na quantidade contida em size.
			// Adciona um novo item a lista chamando a funçao add().
			list.Add()
			if list.Size() != i+1 { // testa se o tamanho da lista é diferente do indice + 1, teste usao para saber se esta no final da lista
				t.Error("%T valor do index n posicao 0 é %d, mas esperavamos que foss %d", list, list.Size(), i+1) // Retorna um erro caso o item adcionado esteja fora do escopo/indice
			}
		}
	}
}

func TestAddOnIndexBeginning(t *testing.T) { // add um item no começo da lista
	defer setupTest()()          // usado para chamar a função setupTest
	for _, list := range lists { // usado para criar uma variavel lista e ao mesmo tempo atribuir um valor na memoria para a mesma
		for i := 0; i < size; i++ {
			// usado para criar e instacia a variavel i que irá ser atualizada baseado na quantidade contida em size.
			list.AddOnIndex(i, 0)   // Adciona um novo item a lista em um determinado index [0].
			if list.Size() != i+1 { // testa se o tamanho da lista é diferente do indice + 1, teste usao para saber se esta no final da lista
				t.Error("%T valor do index na posicao é %d, mas esperavamos que foss %d", list, list.Size(), i+1) // Retorna um erro caso o item adcionado esteja fora do escopo/indice
			}
			val, err := list.get(0) // usado para instancia a variavel err e atribuir o valor da posição 0
			if val != i {           // testa se o valor na posição é igual o valor de val
				t.Error("%T valor do index na posicao é %d, mas esperavamos que foss %d", list, val, i)
			}
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	}
}

func TestAddOnIndexEnd(t *testing.T) { // add um item no começo da lista
	defer setupTest()()          // usado para chamar a função setupTest
	for _, list := range lists { // usado para criar uma variavel lista e ao mesmo tempo atribuir um valor na memoria para a mesma
		for i := 0; i < size; i++ {
			// usado para criar e instacia a variavel i que irá ser atualizada baseado na quantidade contida em size.
			list.AddOnIndex(i, list.Size()) // Adciona um novo item a lista em um determinado index.
			if list.Size() != i+1 {         // testa se o tamanho da lista é diferente do indice + 1, teste usao para saber se esta no final da lista
				t.Error("%T valor do index na posicao é %d, mas esperavamos que foss %d", list, list.Size(), i+1) // Retorna um erro caso o item adcionado esteja fora do escopo/indice
			}
			val, err := list.get(i) // usado para instancia a variavel err e atribuir o valor da posição 0
			if val != i {           // testa se o valor na posição é igual o valor de val
				t.Error("%T valor do index na posicao é %d, mas esperavamos que foss %d", list, val, i)
			}
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	}
}

// Função feita para add item no meio da lista

func TestAddOnIndexMId(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		// preenche a lista com 1's
		for i := 0; i < size; i++ {
			list.Add(1)
		}

		// add -1, 3 times , no index 2
		for i := 0; i < size; i++ {
			list.AddOnIndex(-1, 2)
		}

		for i := 0; i < 8; i++ {
			val, err := list.Get(i)
			if i >= 2 && i < 5 {
				if val != -1 {
					t.Errorf("%T valor no index %d é %d, Mas o esperado seria  -1", list, i, val)
				}
			} else {
				if val != 1 {
					t.Errorf("%T valor na posição %d é %d, mas o esperado era 1", list, i, val)
				}
			}
			if err != nil {
				t.Errorf(err.Error())
			}

		}
	}

}

func TestRemoveOnIndexBeginning(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		// preenche a lista com o valor inserido, seguindo até a posição
		for i := 0; i < size; i++ {
			list.Add(i)
		}

		for i := 0; i < size; i++ {
			// Selecionando o item na primeira posição
			val, err := list.Get(0)
			// teste se o valor Existe na posição
			if val != 1 {
				// Se o erro for diferente de NULL essa será a saída
				t.Errorf("%T valor no index %d é %d, Mas o esperado seria 0", list, val, i)
			}
			if err != nil {
				// saída quando o resultado for diferente de NULL
				t.Errorf(err.Error())
			}
			// Após todos os passos Remove o item da lista
			list.TestRemoveOnIndex(0)
			// testa o tamanho da vista para não remover fora da lista
			if list.Size() != size-i-1 {
				t.Errorf("%T size is %d, but we expected it to be %d", list, list.Size(), size-i-1)
			}
		}
	}

}

func TestRemoveOnIndexEnd(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		// preenche a lista com o valor inserido, seguindo até a posição
		for i := 0; i < size; i++ {
			list.Add(i)
		}
		// Vai percorrer a lista enquanto houver valores
		for i := size - 1; i >= 0; i-- {
			// Selecionando o item na primeira posição
			val, err := list.Get(i)
			if val != i {
				t.Errorf("%T valor no index %d é %d, Mas o esperado seria 0", list, val, i)
			}
			if err != nil {
				t.Errorf(err.Error())
			}
			list.TestRemoveOnIndex(0)
			if list.Size() != i {
				t.Errorf("%T size is %d, but we expected it to be %d", list, list.Size(), i)
			}
		}
	}

}

func TestRemoveOnIndexMid(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		// preenche a lista com o valor inserido, seguindo até a posição
		for i := 0; i < size; i++ {
			list.Add(i)
		}
		// Remove on mid
		list.TestRemoveOnIndex(2)
		for i := 2; i < list.Size(); i++ {
			// Selecionando o item na primeira posição
			val, err := list.Get(i)
			if val != i+1 {
				t.Errorf("%T valor no index %d é %d, Mas o esperado seria 0", list, i, val, i+1)
			}
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	}

}

func TestSet(t *testing.T) {
	defer setupTest()()
	for _, list := range lists {
		//fulfill list with 1's
		for i := 0; i < size; i++ {
			list.Add(1)
		}

		//set -1 on even indexes
		for i := 0; i < size; i++ {
			if i%2 == 0 {
				list.Set(-1, i)
			}
		}

		//check values before index 2 are the same
		for i := 0; i < size; i++ {
			val, err := list.Get(i)
			if i%2 == 0 {
				if val != -1 {
					t.Errorf("%T value on index %d is %d, but we expected it to be -1", list, i, val)
				}
			} else {
				if val == -1 {
					t.Errorf("%T value on index %d is %d, but we expected it to be different from -1", list, i, val)
				}
			}
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	}
}
