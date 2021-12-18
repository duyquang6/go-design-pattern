package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	intCollection := intCollection{arr}
	iter := intCollection.createIterator()
	for iter.hasNext() {
		value, _ := iter.next()
		fmt.Println(value)
	}
}
