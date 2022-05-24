package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Page    int      `json:"page"`
	Entries []string `json:"entries"`
}

func main() {

	str := `{"page":1, "entries":["classA","classB"]}`
	res := response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)

}
