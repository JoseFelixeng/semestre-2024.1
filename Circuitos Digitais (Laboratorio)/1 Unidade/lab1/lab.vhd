library ieee;
use ieee.std_logic_1164.all;

-- Entidade 

entity lab1 is 
    port(
        a, b, c: in bit; 
        ldm11, ldm12,ldm21, ldm22, pt1, pt2, pt3, comutiA1, comutiA2, comutiM1, comutiM2, associA1, associA2, associM1, associM2, dist11, 
		  dist12, dist21, dist22: out bit
    );
end lab1;

architecture main of lab1 is 
begin 

-- Primeira Lei de Morgan 
    -- (A * B)'
    ldm11 <= not(a AND b);
    --(A' +  B')
    ldm12 <= not(a) OR not(b);

-- Segunda lei de Morgan 
    ldm21 <= not(a OR b);
    ldm22 <= not(a) AND not(b);

-- Testando os postulados 
    -- postulado 1 =>  A + 0 = A
    pt1 <= a OR '0';

    -- Postulado 2 => A * 1 = A 
    pt2 <= a AND '1';

    -- Postulado 3 => (A')' = A  	
    pt3 <= not(not(a));

-- Testando as Propriedades de operações Booleanas
    -- Comutatividade - Adição A + B = B + A
    comutiA1 <= a OR b;
    comutiA2 <= b OR a; 

    -- Comutatividade - Multiplicação A * B = B * A
    comutiM1 <= a AND b;
    comutiM2 <= b AND a; 

    -- Associatividade - Adição A + ( B + C )= (A + B) + C
    associA1 <= a OR (b OR c);
    associA2 <= (a OR b) OR c; 

    -- Associatividade - Multiplicação A * ( B * C) = (A * B) * C
    associM1 <= a AND (b AND c);
    associM2 <= (a AND b) AND c; 

    -- Distributiva A * (B + C) = (A * B) + (A * C)
    dist11 <= a AND (b OR c);
    dist12 <= (a AND b) OR (a AND c);

    -- Distributiva A + (B * C) = (A + B) * (A + C)
    dist21 <= a OR (b AND c);
    dist22 <= (a OR b) AND (a OR c);

end architecture main;