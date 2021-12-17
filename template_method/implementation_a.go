package main

import "fmt"

type concreteA struct{ abstractTemplate }

func (c concreteA) step1() {
	fmt.Println("concrete A: step 1")
}

func (c concreteA) step2() {
	fmt.Println("concrete A: step 2")
}

func (c concreteA) step3() {
	fmt.Println("concrete A: step 3")
}
