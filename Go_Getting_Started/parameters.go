package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers!")
	port := 3000
	startWebServer(port, 2)
}

/* func startWebServer(port int, numberOfRetries int) { // is same as below
 * Go assumes that a comma delimited list followed by a data type is all the
 * same data type.
 */
func startWebServer(port, numberOfRetries int) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port ", port)
	fmt.Println("Number of retries ", numberOfRetries)
}
