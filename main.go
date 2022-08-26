package main

import (
	"fmt"
	"net/http"

	"github.com/MrHenri/marketplace-go/routes"
)

func main() {
	routes.LoadRoutes()
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
