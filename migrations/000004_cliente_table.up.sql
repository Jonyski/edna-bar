CREATE TABLE IF NOT EXISTS Cliente (
    id_cliente serial PRIMARY KEY,
    CPF char(11),
    data_nascimento date,
    nome varchar(50) NOT NULL
);
