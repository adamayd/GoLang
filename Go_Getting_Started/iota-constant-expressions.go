package main

import "fmt"

// constants are evaluated at compile time, not run time.  They cannot
// be set to function return values.
const pi = 3.1415

const (
	first  = 1
	second = "second"
)

const (
	third  = iota
	fourth = iota
)

// iota resets between constant blocks.
const (
	fifth = iota + 6
	sixth
	// << is bit shift
	seventh = 2 << iota
)

func main() {
	fmt.Println(pi)

	fmt.Println(first, second)

	fmt.Println(third, fourth)

	fmt.Println(fifth, sixth, seventh)
}
