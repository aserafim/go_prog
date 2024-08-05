package main

import "fmt"

var test int

var p = f()

func f() *int {
	v := 1
	return &v
}

func main() {
	fmt.Println(f())
}
