# Especialização

Processo de definir um conjunto de subclasses de um tipo de entidade(Superclasse).
1. Especialização Total
2. Restrição de Disjunção: **vai para apenas uma das suas restrições**

#### Exemplo 

```sql
Tabela "Funcionario":								
								
cpf         nome   tipo_emprego	     nivel		    grau	    tipo_eng     
1	        Edu     "Técnico"       	            3	     
2	        João	"Secretario"        1           4

```


```sql
Tabela "Secretario":								
								
f_id    nivel  
```

```sql
Tabela "Tecnico":								
								
f_id    grau	      
```

```sql
Tabela "Engenheiro":								
								
f_id     tipo_eng     
```

A tabela Funcionario se relaciona com as tabelas acima por meio de chave estrangeira para modelar de maneira mais adequada o BD.