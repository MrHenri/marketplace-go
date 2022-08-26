package routes

import (
	"net/http"

	"github.com/MrHenri/marketplace-go/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}