# 1 Unidade Circuitos Digitais 
### Semestre 2024.1 Professor (samuel)
#### Anotações Feitas por:  José Felix, baseadas no Livro (Design Digital Frank Vahid) 

## Introdução 
Nos dias atuais há diversas aplicações as quais só foram possíveis graças aos sistemas digitais.

### Digital VS Analógico
* Um sinal digital é aquele que pode assumir um valor de um conjunto finito de valores possívei, a qualquer instante, também pode ser chamado  de sinal `Discreto`.
* Um sinal analógico por sua vez pode ter um valor de um conjunto infinito de valores possíveis e pode ser ser chamado de sinal `contínuo`.
* Um sinal é apenas um fenômeno físico que tem um único valor em cada instante de tempo.

Nos sistemas de computação os sinais digitais mais comuns são aqueles que podem assumir assumir um entre apenas dois valores, como `Ligado` e `Desligado`(representado frequemente como `1` ou  `0`, essa representação de apenas dois valores é conhecida como representação binária.)

* Um sistema digital é aquele que recebe um sinal digital na entrada e gera uma saída digital. 
* Um circuito digital é uma conexão de componentes digitais que juntos constituem um sistema digital.


=> No livro do Frank Vahid o termo `Digital` irá se referir a sistewmas com sinais de valor binário.

Um sinal binário simples é conhecido como `dígito binário` ou `bit` da expressão inglesa "Binary Digit". 
* Uma aplicação bem conhecida dos circuitos digitais são os `microcontroladores`.

* Os circutos digitais são encontrados nas diferentes aplicações dos computadores de própositos gerais são chamados frequentemente de `Sistemas Embarcados`.
Uma vantagem dos sistemas digitais é a resistência a deteriorização ao longo do tempo. A utilização de sinais digitais no lugae de sinais analogicos para guardar audio por exemplo reduz a perda por deteriozação do arquivo pois é mais facil entender para que valor irá pois um sinal digital em um dado espaço de tempo pode ser apenas dois valores ``` 0/1 ```. Além disso outra vantagem para esse tipo de audio é a compreessão dos dados.

#### forma de compressão 

    * 1 - Caso o numero possa a seguinte caracteristica "0000000000"/ "11111111111"
    * 2 - Quando o proximo bit for 0, dessa forma 000000000 ficaria da seguinte forma "00"
    * 3 - Quando o proximo bit for 1, dessa forma 111111111 ficaria da seguinte forma "01"
    * 4 - Se o primeiro de uma amostra for 1, então os proximos 10 bits mostra como o  a mesma realmente é

* Exemplo: 

```number
0000000000 0000000000 0000001111 1111111111
```

* Seria comprimido e ficaria: 

```number
00 00 10000001111 11
```

