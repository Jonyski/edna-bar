CREATE TABLE IF NOT EXISTS aplica (
    id_venda int NOT NULL,
    id_oferta int NOT NULL,

    PRIMARY KEY (id_venda, id_oferta),
    FOREIGN KEY (id_venda) REFERENCES Venda(id_venda) ON DELETE SET NULL,
    FOREIGN KEY (id_oferta) REFERENCES Oferta(id_oferta) ON DELETE SET NULL
);
