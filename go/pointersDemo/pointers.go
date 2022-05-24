package main

import "fmt"

func main() {
	var creature Creature = Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature)
	creature.Reset()
	fmt.Printf("2) %+v\n", creature)
}

type Creature struct {
	Species string
}

func (c Creature) String() string {
	return c.Species
}

func (c *Creature) Reset() {
	c.Species = ""
}
