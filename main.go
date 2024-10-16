package main

import (
	"net/http"

	"github.com/fymorGod/alura/routes"
)

func main() {
	routes.InitRoutes()
	http.ListenAndServe(":8000", nil)
}
