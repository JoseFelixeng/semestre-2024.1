-- Porta  OR
-- POrta OR Generica para ser instanciada mais de uma vez no código
entity PortaOr4is
    port(
        en1, en2, en3  : in bit;
        saida_or       : out bit
        );
end PortaOr4;
   
   
architecture behav of PortaOr4 is
begin
    saida_or <= en1 or en2 or en3;
end architecture behav;