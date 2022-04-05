package main

import "fmt"

type unaryOp func(int) int

func addK(k int) unaryOp {
	return func(a int) int {
		return a + k
	}
}

func main() {
	counter := 1
	func(str string) {
		fmt.Println("Hi", str, "I'm an anonymous function")
	}("Ricky")

	funcVar := func(str string) {
		fmt.Println("Hi", str, "I'm an anonymous function assigned to a variable")
	}

	funcVar("Johnny")

	closure := func(str string) {
		fmt.Println("Hi", str, "I'm a closure")
		for i := 1; i < 5; i++ {
			fmt.Println("Counter incremented to:", counter)
			counter++
		}
	}

	fmt.Println("Counter is ", counter, "before calling closure.")
	closure("Sandy")
	fmt.Println("Counter is ", counter, "after calling closure.")

	add2 := addK(2)
	add3 := addK(3)
	fmt.Println(add2(3)) //print 5
	fmt.Println(add3(4)) //print 7

}
