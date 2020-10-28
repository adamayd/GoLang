package main

type User struct {
  ID int
  FirstName string
  LastName string
}

func main() {
  u1 := User{
    ID: 1,
    FirstName: "Adam",
    LastName: "Ayd",
  }
  u2 := User{
    ID: 2,
    FirstName: "Mike",
    LastName: "Van Sickle",
  }
  // structs are value types and able to be directly compared
  if u1 == u2 {
    println("Same user!")
  } else if u1.FirstName == u2.FirstName {
    println("Similiar users!")
  } else {
    println("Different users!")
  }
}

