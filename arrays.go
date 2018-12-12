package main

import "fmt"
//
//[] means slicing

// input must match the size of the array(not matter pointers or values)
// input is passed by value (copied)
func printArray(arr *[3]int) {
	for i,v := range arr{
		fmt.Println(i,v)
	}
	arr[0] = 100
}

func main() {
	var arr1 [5]int;
	arr2 := [3]int {1,2,3}
	arr3 := [...]int {2, 3, 4}
	for i := range arr3{
		fmt.Println(i)
	}
	for i,v := range arr3{
		fmt.Println(i,v)
	}
	fmt.Println(len(arr1))
	var grid [4][5]int;
	fmt.Println(arr1, arr2, arr3, grid)
	printArray(&arr3)
}
