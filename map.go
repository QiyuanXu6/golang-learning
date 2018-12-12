package main

import "fmt"
// stuct can be key, type can be key, key cannot be func, map, slic
func main() {
	//type: map[K]V, map[K1]map[K2]V
	m := map[string]string {
		"A":"B",
		"B":"C",
	}
	m2 := make(map[string]int) // empty map
	m3 := map[string]string{}
	var m4 map[string]int // nil

	fmt.Println(m2, m3, m4)

	fmt.Println(m["D"]) // print zero value for nonkey
	delete(m, "A")
	val, ok := m["A"] // ok means existence
	fmt.Println(val, ok)
}
