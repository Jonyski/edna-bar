CREATE TABLE IF NOT EXISTS Lote (
    id_lote serial PRIMARY KEY,
    data_fornecimento date NOT NULL,
    id_fornecedor int NOT NULL,

    FOREIGN KEY (id_fornecedor) REFERENCES Fornecedor(id_fornecedor)
);
