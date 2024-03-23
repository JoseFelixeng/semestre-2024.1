-- Porta AND
-- Criando uma porta Generica para usar mais de uma vez
entity PortaAnd2 is
    port(
        en1, en2, en3, en4  : in bit;
        saida_and           : out bit
        );
end PortaAnd2;
   
   
architecture behav of PortaAnd2 is
begin
    saida_and <= en1 and en2 and en3;
end architecture behav;