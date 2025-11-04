package model

import (
	"net/url"
	"strconv"
)

type Produto struct {
	Id int64 `json:"id"`
	Nome string `json:"nome"`
	Categoria string `json:"categoria"`
	Marca string `json:"marca"`
	QntDisponivel uint `json:"qnt_disponivel"`
	QntTotal uint `json:"qnt_total"`
}

type Estrutural struct {
	Produto
}

type Comercial struct {
	Produto
	PrecoVenda float32 `json:"preco_venda"`
}

type ProdutoEstruturalPayload struct {
	Nome string `json:"nome"`
	Categoria string `json:"categoria"`
	Marca string `json:"marca"`
	QntDisponivel uint `json:"qnt_disponivel"`
	QntTotal uint `json:"qnt_total"`
}

type ProdutoComercialPayload struct {
	Nome string `json:"nome"`
	Categoria string `json:"categoria"`
	Marca string `json:"marca"`
	QntDisponivel uint `json:"qnt_disponivel"`
	QntTotal uint `json:"qnt_total"`
	PrecoVenda float32 `json:"preco_venda"`
}


type ComercialFilter struct {
	// Busca, ordenação e paginação
	Nome string
	Categoria string
	Marca string

	Offset uint
	Limit uint
	Sort string

	// Filtragem
	MinQntDisp uint
	MaxQntDisp uint

	MinQntTotal uint
	MaxQntTotal uint

	MinPrecoVenda float32
	MaxPrecoVenda float32
}

type EstruturalFilter struct {
	// Busca, ordenação e paginação
	Nome string
	Categoria string
	Marca string
	Offset uint
	Limit uint
	Sort string

	// Filtragem
	MinQntDisp uint
	MaxQntDisp uint

	MinQntTotal uint
	MaxQntTotal uint
}


func NewComercial(payload ProdutoComercialPayload) *Comercial {
	return &Comercial{
		Produto: Produto{
			Nome: payload.Nome,
			Categoria: payload.Categoria,
			Marca: payload.Marca,
			QntDisponivel: payload.QntDisponivel,
			QntTotal: payload.QntTotal,
		},
		PrecoVenda: payload.PrecoVenda,
	}
}

func NewEstrutural(payload ProdutoEstruturalPayload) *Estrutural {
	return &Estrutural{
		Produto: Produto{
			Nome: payload.Nome,
			Categoria: payload.Categoria,
			Marca: payload.Marca,
			QntDisponivel: payload.QntDisponivel,
			QntTotal: payload.QntTotal,
		},
	}
}

func NewComercialFilter(url *url.URL) *ComercialFilter {
	filter := ComercialFilter {
		Nome: url.Query().Get("nome"),
		Categoria: url.Query().Get("categoria"),
		Marca: url.Query().Get("marca"),
		Sort: url.Query().Get("sort"),
		Offset: 0,
		Limit: 0,
		MinQntDisp: 0,
		MaxQntDisp: 0,
		MinQntTotal: 0,
		MaxQntTotal: 0,
		MinPrecoVenda: 0.0,
		MaxPrecoVenda: 0.0,
	}

	if n, err := strconv.ParseUint(url.Query().Get("offset"), 10, 32); err == nil {
		filter.Offset = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("limit"), 10, 32); err == nil {
		filter.Limit = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("min-qnt-dsp"), 10, 32); err == nil {
		filter.MinQntDisp = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("max-qnt-dsp"), 10, 32); err == nil {
		filter.MaxQntDisp = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("min-qnt-total"), 10, 32); err == nil {
		filter.MinQntTotal = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("max-qnt-total"), 10, 32); err == nil {
		filter.MaxQntTotal = uint(n)
	}
	if n, err := strconv.ParseFloat(url.Query().Get("min-preco-venda"), 32); err == nil {
		filter.MinPrecoVenda = float32(n)
	}
	if n, err := strconv.ParseFloat(url.Query().Get("max-preco-venda"), 32); err == nil {
		filter.MaxPrecoVenda = float32(n)
	}

	return &filter
}

func NewEstruturalFilter(url *url.URL) *EstruturalFilter {
	filter := EstruturalFilter {
		Nome: url.Query().Get("nome"),
		Categoria: url.Query().Get("categoria"),
		Marca: url.Query().Get("marca"),
		Sort: url.Query().Get("sort"),
		Offset: 0,
		Limit: 0,
		MinQntDisp: 0,
		MaxQntDisp: 0,
		MinQntTotal: 0,
		MaxQntTotal: 0,
	}

	if n, err := strconv.ParseUint(url.Query().Get("offset"), 10, 32); err == nil {
		filter.Offset = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("limit"), 10, 32); err == nil {
		filter.Limit = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("min-qnt-dsp"), 10, 32); err == nil {
		filter.MinQntDisp = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("max-qnt-dsp"), 10, 32); err == nil {
		filter.MaxQntDisp = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("min-qnt-total"), 10, 32); err == nil {
		filter.MinQntTotal = uint(n)
	}
	if n, err := strconv.ParseUint(url.Query().Get("max-qnt-total"), 10, 32); err == nil {
		filter.MaxQntTotal = uint(n)
	}

	return &filter
}
