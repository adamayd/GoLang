package main

import "fmt"

func main() {
	fmt.Println("Constants")

	// constants must be assigned a value at declaration
	const pi = 3.1415
	fmt.Println(pi)
	// pi = 3.14 // does not compile

	// uses implicit typing
	const c = 3
	fmt.Println(c + 3)

	fmt.Println(c + 1.2)

	// uses explicit typing
	const co int = 3
	fmt.Println(co + 3)

	// type mismatch
	// fmt.Println(co + 1.2)
	fmt.Println(float32(co) + 1.2)
}
