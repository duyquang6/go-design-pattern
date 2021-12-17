package main

import "fmt"

type concreteB struct{}

func (c concreteB) step1() {
	fmt.Println("concreteB: step 1")
}

func (c concreteB) step2() {
	fmt.Println("concreteB: step 2")
}

func (c concreteB) step3() {
	fmt.Println("concreteB: step 3")
}
