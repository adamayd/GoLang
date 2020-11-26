package main

import (
	"net/http"

	"github.com/adamayd/GoLang/Create_Web_Services_with_Go/working/inventory-service/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
