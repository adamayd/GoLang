package main

import (
	"log"
	"net/http"
	"webservice/database"
	"webservice/product"
	"webservice/receipt"

	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	product.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
