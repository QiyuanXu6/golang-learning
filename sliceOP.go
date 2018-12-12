package main

import "fmt"

func main() {
	var s []int;// this is a slice, zero as nil
	var s2 [16]int;// this is an arr
	s3 := make([]int, 16) // this is a slice
	s4 := make([]int, 16, 30) // this is a slice
	fmt.Println(s, s2, s3, s4)

	s5 := []int{1, 2, 3}
	copy(s3, s5)
	fmt.Println(s3)

	// delete
	s5 = append(s5[:2], s5[3:]...)
}
