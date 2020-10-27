package main

import (
  "github.com/adamayd/webservice/models"
  "fmt"
)

func main() {
  u := models.User{
    ID: 2,
    FirstName: "Mike",
    LastName: "Van Sickle",
  }
  fmt.Println(u)
}
