package produto

import (
	"context"
	"database/sql"
	"edna/internal/model"
	"edna/internal/types"
	"fmt"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAllComercial(ctx context.Context, filter *model.ComercialFilter) ([]model.Comercial, error) {
	query := "SELECT id_produto, nome, categoria, marca, quantidade_disponivel, quantidade_total, preco_venda FROM ProdutoComercial JOIN Produto USING (id_produto)"
	rows, err := s.queryRowsWithComercialFilter(ctx, query, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	produtos := make([]model.Comercial, 0)
	for rows.Next() {
		c := model.Comercial{}
		err = rows.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca, &c.QntDisponivel, &c.QntTotal, &c.PrecoVenda)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, types.ErrNotFound
			}
			log.Printf("Error scanning row: %v", err)
			return nil, types.ErrInternalServer
		}
		produtos = append(produtos, c)
	}

	return produtos, nil
}


func (s *Store) GetAllEstrutural(ctx context.Context, filter *model.EstruturalFilter) ([]model.Estrutural, error) {
	query := "SELECT id_produto, nome, categoria, marca, quantidade_disponivel, quantidade_total FROM ProdutoEstrutural JOIN Produto USING (id_produto)"
	rows, err := s.queryRowWithEstruturalFilter(ctx, query, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	produtos := make([]model.Estrutural, 0)
	for rows.Next() {
		c := model.Estrutural{}
		err = rows.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca, &c.QntDisponivel, &c.QntTotal)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, types.ErrNotFound
			}
			log.Printf("Error scanning row: %v", err)
			return nil, types.ErrInternalServer
		}
		produtos = append(produtos, c)
	}

	return produtos, nil
}

func (s *Store) queryRowsWithComercialFilter(ctx context.Context, query string, filter *model.ComercialFilter) (*sql.Rows, error) {
	var filterValues []any

	// Adicionamos o WHERE aqui, se não houver nome usamos true (anulando a condição)
	// Se houver, ótimo, e podemos encadear com as outras condições
	query += " WHERE"
	if filter.Nome != "" {
		filterValues = append(filterValues, filter.Nome)
		query += " nome LIKE '%%' || $1 ||'%'"
	} else {
		// Adicionamos true para adicionar o
		query += " true"
	}
	if filter.Categoria != "" {
		filterValues = append(filterValues, filter.Categoria)
		query += fmt.Sprintf(" AND categoria LIKE '%%' || $%d ||'%%'", len(filterValues))
	}
	if filter.Marca != "" {
		filterValues = append(filterValues, filter.Marca)
		query += fmt.Sprintf(" AND marca LIKE '%%' || $%d ||'%%'", len(filterValues))
	}

	if filter.MinQntDisp > 0 {
		filterValues = append(filterValues, filter.MinQntDisp)
		query += fmt.Sprintf(" AND quantidade_disponivel >= $%d", len(filterValues))
	}
	if filter.MaxQntDisp > 0 {
		filterValues = append(filterValues, filter.MaxQntDisp)
		query += fmt.Sprintf(" AND quantidade_disponivel <= $%d", len(filterValues))
	}

	if filter.MinQntTotal > 0 {
		filterValues = append(filterValues, filter.MinQntTotal)
		query += fmt.Sprintf(" AND quantidade_total >= $%d", len(filterValues))
	}

	if filter.MaxQntTotal > 0 {
		filterValues = append(filterValues, filter.MaxQntTotal)
		query += fmt.Sprintf(" AND quantidade_total <= $%d", len(filterValues))
	}

	if filter.MinPrecoVenda > 0.0 {
		filterValues = append(filterValues, filter.MinPrecoVenda)
		query += fmt.Sprintf(" AND preco_venda >= $%d", len(filterValues))
	}

	if filter.MaxPrecoVenda > 0.0 {
		filterValues = append(filterValues, filter.MaxPrecoVenda)
		query += fmt.Sprintf(" AND preco_venda <= $%d", len(filterValues))
	}

	if filter.Offset > 0 {
		filterValues = append(filterValues, filter.Offset)
		query += fmt.Sprintf(" OFFSET $%d", len(filterValues))
	}
	if filter.Limit > 0 {
		filterValues = append(filterValues, filter.Limit)
		query += fmt.Sprintf(" LIMIT $%d", len(filterValues))
	}

	if filter.Sort != "" {
		filterValues = append(filterValues, filter.Sort)
		query += fmt.Sprintf(" ORDER BY $%d", len(filterValues))
	}

	return s.db.QueryContext(ctx, query, filterValues...)
}

func (s *Store) queryRowWithEstruturalFilter(ctx context.Context, query string, filter *model.EstruturalFilter) (*sql.Rows, error) {
	var filterValues []any

	// Adicionamos o WHERE aqui, se não houver nome usamos true (anulando a condição)
	// Se houver, ótimo, e podemos encadear com as outras condições
	query += " WHERE"
	if filter.Nome != "" {
		filterValues = append(filterValues, filter.Nome)
		query += " nome LIKE '%%' || $1 ||'%'"
	} else {
		// Adicionamos true para adicionar o
		query += " true"
	}
	if filter.Categoria != "" {
		filterValues = append(filterValues, filter.Categoria)
		query += fmt.Sprintf(" AND categoria LIKE '%%' || $%d ||'%%'", len(filterValues))
	}
	if filter.Marca != "" {
		filterValues = append(filterValues, filter.Marca)
		query += fmt.Sprintf(" AND marca LIKE '%%' || $%d ||'%%'", len(filterValues))
	}

	if filter.MinQntDisp > 0 {
		filterValues = append(filterValues, filter.MinQntDisp)
		query += fmt.Sprintf(" AND quantidade_disponivel >= $%d", len(filterValues))
	}
	if filter.MaxQntDisp > 0 {
		filterValues = append(filterValues, filter.MaxQntDisp)
		query += fmt.Sprintf(" AND quantidade_disponivel <= $%d", len(filterValues))
	}

	if filter.MinQntTotal > 0 {
		filterValues = append(filterValues, filter.MinQntTotal)
		query += fmt.Sprintf(" AND quantidade_total >= $%d", len(filterValues))
	}

	if filter.MaxQntTotal > 0 {
		filterValues = append(filterValues, filter.MaxQntTotal)
		query += fmt.Sprintf(" AND quantidade_total <= $%d", len(filterValues))
	}

	if filter.Sort != "" {
		filterValues = append(filterValues, filter.Sort)
		query += fmt.Sprintf(" ORDER BY $%d", len(filterValues))
	}

	if filter.Offset > 0 {
		filterValues = append(filterValues, filter.Offset)
		query += fmt.Sprintf(" OFFSET $%d", len(filterValues))
	}
	if filter.Limit > 0 {
		filterValues = append(filterValues, filter.Limit)
		query += fmt.Sprintf(" LIMIT $%d", len(filterValues))
	}

	return s.db.QueryContext(ctx, query, filterValues...)
}

func (s *Store) CreateComercial(ctx context.Context, props *model.Comercial) error {
	// Executa uma transação (desfaz insercao em caso de erro) pois precisamos fazer 2 insercoes
	t, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer t.Rollback()

	query1 := "INSERT INTO Produto (nome, categoria, marca, quantidade_disponivel, quantidade_total) VALUES ($1, $2, $3, $4, $5) RETURNING id_produto;"
	query2 := "INSERT INTO ProdutoComercial (id_produto, preco_venda) VALUES ($1, $2) RETURNING id_produto;"

	r1 := t.QueryRowContext(ctx, query1, props.Nome, props.Categoria, props.Marca, props.QntDisponivel, props.QntTotal)
	var idProduto int64
	err = r1.Scan(&idProduto)
	if err != nil {
		return err
	}

	_, err = t.ExecContext(ctx, query2, idProduto, props.PrecoVenda)
	if err != nil {
		return err
	}

	err = t.Commit()
	if err != nil {
		return err
	}
	// Atualiza o id do produto
	props.Id = idProduto
	return nil
}

func (s *Store) CreateEstrutural(ctx context.Context, props *model.Estrutural) error {
	// Executa uma transação (desfaz insercao em caso de erro) pois precisamos fazer 2 insercoes
	t, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer t.Rollback()

	query1 := "INSERT INTO Produto (nome, categoria, marca, quantidade_disponivel, quantidade_total) VALUES ($1, $2, $3, $4, $5) RETURNING id_produto;"
	query2 := "INSERT INTO ProdutoEstrutural (id_produto) VALUES ($1);"

	r1 := t.QueryRowContext(ctx, query1, props.Nome, props.Categoria, props.Marca, props.QntDisponivel, props.QntTotal)
	var idProduto int64
		err = r1.Scan(&idProduto)
		if err != nil {
			return err
		}

	_, err = t.ExecContext(ctx, query2, idProduto)
	if err != nil {
		return err
	}

	err = t.Commit()
	if err != nil {
		return err
	}

	// Atualiza o id do produto
	props.Id = idProduto
	return nil
}
func (s *Store) UpdateComercial(ctx context.Context, id int64, props *model.Comercial) error {
	t, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer t.Rollback()
	query := "UPDATE Produto SET nome = $1, categoria = $2, marca = $3, quantidade_disponivel = $4, quantidade_total = $5 WHERE id_produto = $6;"
	res, err := t.ExecContext(ctx, query, props.Nome, props.Categoria, props.Marca, props.QntDisponivel, props.QntTotal, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return types.ErrNotFound
	}

	query2 := "UPDATE ProdutoComercial SET preco_venda = $1 WHERE id_produto = $2;"
	res, err = t.ExecContext(ctx, query2, props.PrecoVenda, id)
	if err != nil {
		return err
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return types.ErrNotFound
	}

	t.Commit()
	props.Id = id
	return nil
}

func (s *Store) UpdateEstrutural(ctx context.Context, id int64, props *model.Estrutural) error {
	query := "UPDATE Produto SET nome = $1, categoria = $2, marca = $3, quantidade_disponivel = $4, quantidade_total = $5 WHERE id_produto = $6;"

	res, err := s.db.ExecContext(ctx, query, props.Nome, props.Categoria, props.Marca, props.QntDisponivel, props.QntTotal, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return types.ErrNotFound
	}
	props.Id = id
	return nil
}

func (s *Store) GetComercialByID(ctx context.Context, id int64) (*model.Comercial, error) {
	query := "SELECT id_produto, nome, categoria, marca, quantidade_disponivel, quantidade_total, preco_venda FROM ProdutoComercial JOIN Produto USING (id_produto) WHERE id_produto = $1"
	row := s.db.QueryRowContext(ctx, query, id)
	c := model.Comercial{}
	err := row.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca, &c.QntDisponivel, &c.QntTotal, &c.PrecoVenda)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, types.ErrNotFound
		case sql.ErrConnDone:
			log.Println("[WARN] Connection closed.")
			return nil, types.ErrInternalServer
		default:
			log.Println("[ERROR] Unexpected error:", err)
			return nil, err
		}
	}
	return &c, nil
}

func (s *Store) GetEstruturalByID(ctx context.Context, id int64) (*model.Estrutural, error) {
	query := "SELECT id_produto, nome, categoria, marca, quantidade_disponivel, quantidade_total FROM ProdutoEstrutural JOIN Produto USING (id_produto) WHERE id_produto = $1"
	row := s.db.QueryRowContext(ctx, query, id)
	c := model.Estrutural{}
	err := row.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca, &c.QntDisponivel, &c.QntTotal)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *Store) Delete(ctx context.Context, id int64) error {
	// Derivadas do produto serão apagadas automaticamente por conta do condicional CASCADE
	query := "DELETE FROM Produto WHERE id_produto = $1"
	_, err := s.db.ExecContext(ctx, query, id)
	return err
}
