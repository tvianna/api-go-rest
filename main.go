package main

import (
	"github.com/tvianna/api-go-rest/models"
	"github.com/tvianna/api-go-rest/routes"

	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},

	}

	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}