
CREATE TABLE IF NOT EXISTS contem_item_lote (
    id_lote int NOT NULL,
    id_produto int NOT NULL,
    preco_unitario decimal(6, 2) NOT NULL,
    tipo_produto tipo_de_produto NOT NULL,
    quantidade_comprada int NOT NULL,
    quantidade_utilizavel int,
    quantidade_manutencao int,
    quantidade_quebrada int,
    validade date,
    quantidade_disponivel int,
    quantidade_armazenada int,
    quantidade_estragada int,
    PRIMARY KEY (id_lote, id_produto),

    FOREIGN KEY (id_lote) REFERENCES Lote(id_lote) ON DELETE CASCADE,
    FOREIGN KEY (id_produto) REFERENCES Produto(id_produto) ON DELETE CASCADE
);
