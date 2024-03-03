# 1 Unidade 
### Anotações da disciplina de Banco de Dados 
#### author: José Felix

## Conceitos 

Um Banco de dados é uma coleção de dados relacionados com dados queremos dizer fatos conhecidos que podem ser armazenados e possuem significado implícito.
`Obs: Uma variedade aleatoria de dados não pode ser corretamente chamada de banco de dados.`
* Possui próposito específico
* Tem estrutura lógica que confere significado aos dados.
* Projetados, constuídos e populados com dados para uma finalidade específica.
* Podem ter qualquer tamanho e complexidade.
    => IRPF: 100 * 10^6 * 400 * 5 Caracteres(bytes) => 800GB a cada 3 anos (exemplo de uma empresa).

## SGDB 

Um banco de dados pode ser gerado e mantido manualmente, ou pode ser computadorizado. Um SGDB é uma coleção de programas que permite aos usuários criar e manter um banco de dados.
Definir um banco de dados esta atrelado a especificar os tipos, estruturas  e restrições dos dados a serem armazenados. A definição propriamente dita é armazenada no SGBD na forma de catalogo, chamado `metadados`.

    Manipulação ============>   Consulta
                ============>   Atualização 


    Vantagens   ============> Abstração de dados
                ============> Suporte a visões dos dados
                ============> Controle de concorrência


#### Caracteristicas Avançadas
* Controle de Redundância
* Normalização
* Controle de Acesso
* Estrutura de armazenamento e técnica de pesquisa para o processamento eficiente de consulta.
* Restrições de integridade


** Chave Primária: É uma `key` usada para identificar o objeto da tabela ao qual esta sendo editada.
** Chave Secundária: É a utilização/LINk de uma outra `key` de uma outra tabela de forma a usar as suas informações.












Exemplo: 



| name | email | present | receiveCertificate | course |
| --- | --- | --- | --- | --- |
| Chaiana Hermes | mailto:chaiana_hermes@yahoo.com.br | true | false | Bootcamp React |
| Chaiana Hermes | mailto:chaiana_hermes@yahoo.com.br | true | false | Bootcamp React |




