# Arquitetura Hexagonal
## Ports and Adapters

![Desenho da arquitetura hexagonal](https://imgs.search.brave.com/C5UDVDtbCjnqlvEDIzsxHX-y6MIuyOyyMGR29Y-jpps/rs:fit:860:0:0/g:ce/aHR0cHM6Ly9zcGlu/LmF0b21pY29iamVj/dC5jb20vd3AtY29u/dGVudC91cGxvYWRz/L3BvcnRzX2FuZF9h/ZGFwdGVyc19hcmNo/aXRlY3R1cmUuanBn "Ports and Adapters")


Essa arquitetura serve para que o core da aplicação seja isolado de ferramentas externas, como libs e frameworks e também para que os componentes não possuam um alto acoplamento entre si.