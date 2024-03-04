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


# Exemplo:
* Grupo 2: Registro de Notas de Alunos em Disciplinas Universitárias
    ### Disciplinas e Alunos:
    A planilha deve conter informações sobre as disciplinas oferecidas pelo departamento, incluindo o código da disciplina, nome do curso, semestre, professor responsável, entre outros detalhes. Deve haver também uma lista dos alunos matriculados em cada disciplina, contendo seus nomes, números de matrícula e outras informações relevantes.

    ### Registro de Notas:
    Os professores devem poder registrar as notas dos alunos em cada avaliação (provas, trabalhos, participação em sala de aula, etc.). As notas devem ser atribuídas individualmente para cada aluno e para cada atividade avaliativa. A planilha deve permitir a flexibilidade para incluir diferentes tipos de avaliações e seus pesos correspondentes.

    ### Cálculo de Médias:
    A planilha deve calcular automaticamente as médias parciais e finais dos alunos com base nas notas registradas. As médias podem ser ponderadas de acordo com os pesos atribuídos a cada atividade avaliativa. 

    ### Feedback aos Alunos:
    Além de registrar as notas, os professores podem adicionar comentários ou feedback sobre o desempenho de cada aluno. Isso permite uma comunicação clara e transparente sobre o progresso acadêmico de cada estudante. 

    ### Relatórios de Desempenho: 
    A planilha deve gerar relatórios automáticos sobre o desempenho dos alunos em cada disciplina. Os relatórios podem incluir as notas de cada aluno, médias parciais e finais, gráficos de desempenho, entre outras informações relevantes.

    ### Acompanhamento do Progresso:
    Os professores e coordenadores acadêmicos devem poder acompanhar o progresso dos alunos ao longo do semestre. Isso inclui identificar alunos em risco, monitorar tendências de desempenho e tomar medidas para ajudar os estudantes que estão com dificuldades.


#### Resolução feita pelo grupo:


```sql
Tabela "Aluno":		
		
id_aluno	id_curso	Nome_Aluno
1	        101	        João
2	        102	        Maria
3	        103	        Pedro
```


```sql
Tabela "Desempenho":								
								
id_desempenho	id_curso	id_aluno		id_professor	Nota1	     Nota2	Nota3	Média	Feedback
1	                101	        1	        201	        8.5	     8.5	 8.5	 8.5	    Bom
2	                102	        2	        202	        9.0	     9.0	 9.0	 9.0	    Otimo
3	                103	        3	        203	        7.8	     7.8	 7.8	 7.8	    Bom 
4	                104	        4	        204	        6.5	     6.5	 6.5	 6.5	    Precisa estudar mais
```

```sql
Tabela "Professor":	
	
Id_professor	Nome_Professor
201	            Silva
202         	    Santos
203	            Oliveira
204	            Pereira
205	            Costa
```

```sql
Tabela "Relação":	
	
id_curso	id_disciplina
101	        101
102	        102
103	        103
104	        104
105	        105
```

```sql
Tabela "curso":	
	
id_curso	Nome_Curso
101	        Engenharia
102	        Medicina
103	        Direito
104	        Psicologia
105	        Arquitetura
```

```sql

Tabela "leciona":			
			
id_disciplina	Id_professor	ano	      semestre
101	            201		        2024        1
102	            202		        2024        1
103	            203		        2024        1
104	            204		        2024        1
101	            205	                2024        1
```

```sql

Tabela "disciplinas":		
		
id_disciplina	Nome_disciplina	             semestre
101         Engenharia Constitucional	        1
102	        Anatomia Humana	           	        2
103	        Direito Constitucional	    	    3
104	        Psicologia Clínica	                4
105	        Design de Interiores	   	        5
```

Podemos vê como ficou a organização das planilhas de forma a melhor adaptar/modelar para o problema proposto.


### Anotações aula presencial 
Relação = Tabela 
Banco de dados relacional = Tabela 
Conceitual => Lógico => Físico


Lógica: está atrelado a olha o diagrama e vê as melhores formas de relacionar

Meio Físico está relacionado a parte de mais baixo nível que diz respeito a otimização 

Exemplo:
Funcionario 
Nome              Salário               dept
Eduardo            X                       DCA
João               Y                       DCA
Maria              Z                       IMP
Ana                A                       DIMAP

