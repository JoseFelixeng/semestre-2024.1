# Curso sobre VHDL  

O vhdl é uma linguagem a nivel de maquina que pode ser usada para auxiliar na interração com o hardware

Estrutura do código VHDL:
1. Biblioteca e pacotes
2. Entidade
3. Arquitetura 

Exemplo porta AND:

```vhdl
--Exemplo porta AND: 
-- Autor: José Felix R. Anselmo
-- Data: 14/03/2024 

-- Entidade 
-- entity nome_do_arquivo is
-- port é utilizado para mostrar a interface das portas
entity and_gate is 
    port(
        -- Variaveis de entrada
        a, b : in  bit;
        -- Variaveis de sáida
        z    : out bit
    );
-- finalizando a entidade
end entity and_gate;

-- Arquitetura 
architecture main of and_gate is 
    -- Definir objetos 
-- Iniciando o código
begin 
    z <= a and b;

--fechando a arquitetura

end architecture main;

```