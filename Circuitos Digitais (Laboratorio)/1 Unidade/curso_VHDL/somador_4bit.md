# exemplo somador 4 bits no VHDL 

```vhdl
--Exemplo porta AND: 
-- Autor: José Felix R. Anselmo
-- Data: 14/03/2024 

-- A palavra reservada integer se refere a variaveis inteiras.
-- para esse tipo de variavel tem que ser criada um range que é o tamanho pelo qual a mesma vai ter.

entity somador is 
    port(
        a, b : in integer range 0 to 15;
        z    : in integer range 0 to 15 
    );
end somador;

architecture main od somador is 
begin 
    z <= a + b;

end main; 
-- fim do somador
```