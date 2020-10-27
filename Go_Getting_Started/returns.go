package main

import (
  "fmt"
  "errors"
)

func main() {
	port := 3000
  // _ is the write only variable when you don't need a variable
  _, err := startWebServer(port)
  fmt.Println(err)
}

/* func startWebServer(port int, numberOfRetries int) { // is same as below
 * Go assumes that a comma delimited list followed by a data type is all the
 * same data type.
 */
func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
  return port, errors.New("Something went wrong.")
}
