package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

type Girl struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Where      string `json:"where"`
	Is_married bool   `json:"is_married"`
}

func main() {
	g := Girl{"Satori", 16, "F", "Oriental spirit ball", false}
	ret, err := json.MarshalIndent(g, "", " ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("**********************")
		fmt.Println("Encoding using json")
		fmt.Println("**********************")
		fmt.Println(string(ret))
	}

	// deserialization
	g2 := Girl{}
	err = json.Unmarshal(ret, &g2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(g2)
	fmt.Println(g2.Name, g2.Age)

	//using gob
	g3 := Girl{"Satori", 16, "F", "Oriental spirit ball", false}
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	if err := enc.Encode(g3); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("**********************")
		fmt.Println("Encoding using gob")
		fmt.Println("**********************")
		fmt.Println(buf.String())
	}

	var g3decoded = Girl{}
	//decodign using gob
	dec := gob.NewDecoder(bytes.NewBuffer(buf.Bytes()))
	err = dec.Decode(&g3decoded)
	if err == nil {
		fmt.Println("**********************")
		fmt.Println("Decoded using gob")
		fmt.Println("**********************")
		fmt.Println(g3decoded)
	} else {
		fmt.Println(err)
	}
}
