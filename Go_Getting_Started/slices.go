package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]

	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4, 42, 27)
	fmt.Println(slice2)

	slice3 := slice2[1:]
	slice4 := slice2[:2]
	slice5 := slice2[1:2]

	fmt.Println(slice3, slice4, slice5)
}
