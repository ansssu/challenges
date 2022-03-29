package main

import "fmt"

var c, python, java bool
var i, j int = 1, 2

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b, i, j, c, python, java)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
