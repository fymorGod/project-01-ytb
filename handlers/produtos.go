package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/fymorGod/alura/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Error na conversão do preco:", err)
		}
		qtdConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Error na conversão na quantidade:", err)
		}
		models.CriarNovoProduto(nome, descricao, precoConvertido, qtdConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletarProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditarProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertida, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error na conversão do id:", err)
		}
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Error na conversão do preco:", err)
		}
		qtdConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Error na conversão na quantidade:", err)
		}
		models.UpdateProduto(idConvertida, nome, descricao, precoConvertido, qtdConvertida)
	}
	http.Redirect(w, r, "/", 301)
}
