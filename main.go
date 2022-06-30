package main

import (
	"github.com/tvianna/api-go-rest/database"
	"github.com/tvianna/api-go-rest/models"
	"github.com/tvianna/api-go-rest/routes"

	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}