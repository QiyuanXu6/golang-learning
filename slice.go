package main

import "fmt"
// slice: ptr, len, cap
func main() {
	arr1 := [...]int {1,2,3}
	// extending: only in back direction
	s1 := arr1[1:2]
	s2 := s1[0:2] // based on the original arr, cap(arr)
	fmt.Println(s2)

	// append: brucely change the value in the tail, but not change original capacity
	// if the length is not enough create another arr to hold the view
	// arr contains locationa and length
	s3 := append(s1,10)
	s4 := append(s3, 11)
	fmt.Println(arr1, s1, s3, s4)

}
