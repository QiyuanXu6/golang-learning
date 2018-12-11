package main

import "fmt"



/// enum
func enums() {
	const (
		b = 1 << (10* iota)
		kb
		mb
	)
	const(
		cpp = iota
		_
		java
		c
	)
	fmt.Println(cpp, java, c)
}

// const: var, kind of like a text replacement->no need to cast if didn't assign type
const (
	filename = "OK.txt"
)


// type: bool string, (u)int, int8, int16, int32, int64, uintprt
//       byte(8bits), rune(char(s))(4bytes)
// 		 float32, float64, complex64, complex128(contains two float of 64bits)

// no implicit type casting! float64(v)

// var name first, type second
// var have default values
// var have type deduction

// no global, but package level variables
var ( // don't need to repeating "var"
	aaa = 3
	eee = 4
	com1 = 1 + 1i
)
func variableZeroVal() {
	var a int
	var s string
	var b, ss, bo = 1, "a", true
	ee, sss := "ee", "sss" // colon means define, try not to use "var", cannot use ":=" outside function
	fmt.Println(a, s, b, ss, bo, ee, sss)
}

func main() {
	fmt.Println("hello word")
	variableZeroVal()
	enums()
}
