package main

import "fmt"

func main() {
	// use * to assign pointer
	var firstName *string = new(string)

	// de-referencing operation by preceding the variable name with *
	*firstName = "Adam"

	fmt.Println(firstName)
	//apply de-referencing operation again to get the value
	fmt.Println(*firstName)

	otherName := "Angela"
	fmt.Println(otherName)

	// Address-of operator is the &
	ptr := &otherName
	fmt.Println(ptr, *ptr)

	otherName = "Angie"
	fmt.Println(ptr, *ptr)
}
