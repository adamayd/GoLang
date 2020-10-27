package main

import "fmt"

func main() {
  type user struct {
    ID int
    FirstName string
    LastName string
  }

  var u user
  u.ID = 1
  u.FirstName = "Adam"
  u.LastName = "Ayd" 
  fmt.Println(u)
  fmt.Println(u.FirstName)

  u2 := user{ ID: 1, FirstName: "Adam", LastName: "Ayd"}
  fmt.Println(u2)

  u3 := user{
    ID: 1,
    FirstName: "Adam",
    LastName: "Ayd",
  }
  fmt.Println(u3)
}
