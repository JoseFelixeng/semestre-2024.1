package list

// Escopo para construção de uma lista.

type Ilist {
	Add(value int) // Adcionar algo a lista 
	AddOnIndex(value int, index int) error // Adciona algo em um lugar desejado
	RemoveOnIndex(value int) error // Revome algo da lista em qualquer posição
	Get(index int) (int, error) // pegar algo/visualiza/dar uma saída da  lista
	Set(value int, index int) error // pegar algo na lista e muda as suas informações/configurações
	Size() int // retorna o tamanho da lista 
}