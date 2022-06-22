package main

import "fmt"

func main() {
	str := join("Hello", "", "World")
	fmt.Println(str)
}

func join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}
