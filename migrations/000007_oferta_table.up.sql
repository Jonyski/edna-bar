CREATE TABLE IF NOT EXISTS Oferta (
    id_oferta serial PRIMARY KEY,
    nome varchar(50) NOT NULL,
    data_criacao date NOT NULL DEFAULT CURRENT_DATE,
    data_inicio date,
    data_fim date,
    valor_fixo decimal(6, 2),
    percentual_desconto int
);
