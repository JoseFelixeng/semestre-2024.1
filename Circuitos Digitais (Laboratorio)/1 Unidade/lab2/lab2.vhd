-- Codigo instanciando as portas


entity lab2_c is
    port(
        A, B, C : in bit;
        S1, S2 : out bit
        );
end lab2_c;
   
   
architecture behav of lab2_c is
    signal aux1: bit;
    signal aux2: bit;
    signal aux3: bit;
    signal aux4: bit;
    signal aux5: bit;
    signal aux6: bit;
    signal aux7: bit;
   
   
   
    component PortaAnd2 is
        port(
            en1, en2     : in bit;
            saida_and : out bit
            );
    end component;
       
    component PortaAnd3 is
        port(
            en1, en2, en3 : in bit;
            saida_and : out bit
        );
    end component;
   
   
    component PortaOR4 is
        port(
            en1, en2, en3, en4 : in bit;
            saida_or           : out bit
            );
        end component;
       
    component PortaOR3 is
        port(
            en1, en2, en3 : in bit;
            saida_or      : out bit
            );
    end component;
   
   
begin
    -- Código do primeiro circuito
    u6 : PortaAnd2 port map(en1 => A, en2 => B, saida_and => aux5);
    u7 : PortaAnd2 port map(en1 => B, en2 => C, saida_and => aux6);
    u8 : PortaAnd2 port map(en1 => A, en2 => C, saida_and => aux7);
       
    u9 : PortaOR3 port map(en1 => aux5, en2 => aux6, en3 => aux7, saida_or => S1);
         
    -- Código do segundo circuito
    u1 : PortaAnd3 port map(en1 => not(A),  en2 => not(B), en3 => C, saida_and => aux1);
    u2 : PortaAnd3 port map(en1 => A,       en2 => not(B), en3 => not(C), saida_and => aux2);
    u3 : PortaAnd3 port map(en1 => not(A),  en2 => B,       en3 => not(C), saida_and => aux3);
    u4 : PortaAnd3 port map(en1 => A,       en2 => B,       en3 => C,       saida_and => aux4);
       
    u5 : PortaOR4 port map(en1 => aux1, en2 => aux2, en3 => aux3, en4 => aux4, saida_or => S2);
       
   
end architecture behav;
