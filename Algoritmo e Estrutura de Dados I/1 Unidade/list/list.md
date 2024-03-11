# Listas

    A lista é uma ferramenta muito importante no dia a dia, que auxilia a gerenciar e ao mesmo tempo organizar uma quantidade infinita de objetos entre outras coisas. para a programação listas são ordem de vetores especificos que podem variar de linguagem para linguagem a sua sintaxe e algumas funcionalidade mais de maneira geral deverão funcionar igual as listas de notas.
    No go podemos podemos instanciar uma lista da seguinte forma.
```go 
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
        
```
