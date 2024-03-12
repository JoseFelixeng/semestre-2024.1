## Aula 3 
### TAD vs ED 
1. TAD (Tipos Abstratos de Dados)
    * Nome
    * Operações
    * Genérico 
    * Lista
2. ED (Estrutura de dados)
    * define
    * estrutura para armazenar e organizar dados

#### Interface Lista 
    1. Adcionar (elemento)
    2. inserir (elemento, index)
    3. remover (elemento)
    4. remover(elemento, index)
    5. set (elemento, index )
    6. size ()
**OBS:** Um array é armazenado sequeêncialmente na RAM.

#### ArrayList

```| 2 | 4 | 6 | 8 | 10 | 12 | ¨ | n | ```
| #Negativos                |           #Positivos  | 
|---|---|
| Redimensionar pode ser custoso O(n) |   Capacidade de acesar qualquer posição do array em O(1) |
| Tende a possuí espaços ociosos  |  End + index sizeof(type)  |


#### Análise de Complexidade:
    1- O(1): Instãntaneo/ constante o tempo não é impactado pelo tamanho da entrada(vetor)
    2- O(n): Tempo é impactado pelo tamanho da entrada de forma que quanto maior a quantidade de dados da entrada mmais lento será a analise

Discutir o desempenho sem recorrer o tempo absoluto. Contamos a quantidade de operações primiticas executadas
* retornar, aritmética, atribuição, definição de variável;
* O(1), O(n), Ω(1), Ω(n),  θ(n),  θ(n^2)
    

