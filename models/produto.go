package models

import "github.com/fymorGod/alura/db"

type Produtos struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produtos {
	db := db.ConnectToDb()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produtos{}
	produtos := []Produtos{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectToDb()

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}
