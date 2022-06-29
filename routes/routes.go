package routes

import (
	"github.com/tvianna/api-go-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest(){
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}