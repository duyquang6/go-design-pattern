package main

import "fmt"

func main() {
	builderA := &builderA{}
	director := director{builderA}
	fmt.Println(director.make("A"))

	builderB := &builderB{}
	director.changeBuilder(builderB)
	fmt.Println(director.make("B"))
}
