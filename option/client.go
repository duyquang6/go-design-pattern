package main

import "fmt"

func main() {
	l := NewLongStructDefinition(1, "abc", true, WithOptA(true), WithOptB("xyz"))

	fmt.Println(l)
}
