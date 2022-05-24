package main

import "fmt"

type CofeeMachine struct {
	NumberOfCoffeeBeans int
}

func (cm *CofeeMachine) SetNumberOfCoffeBeans(n int) {
	cm.NumberOfCoffeeBeans = n
}

func NewCofeeMachine() *CofeeMachine {
	return &CofeeMachine{}
}

func main() {
	cm := NewCofeeMachine()
	cm.SetNumberOfCoffeBeans(100)
	fmt.Printf("The cofee machinehas %d beans\n", cm.NumberOfCoffeeBeans)
}
