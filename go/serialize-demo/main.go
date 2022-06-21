package main

import (
	"encoding/json"
	"fmt"
)

type Girl struct {
	Name       string
	Age        int
	Gender     string
	Where      string
	Is_married bool
}

func main() {
	g := Girl{"Satori", 16, "F", "Oriental spirit ball", false}
	ret, err := json.MarshalIndent(g, "", " ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(ret))
	}
}
