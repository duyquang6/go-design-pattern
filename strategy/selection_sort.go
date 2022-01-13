package main

import "fmt"

var _ unstableSort = (*selectionSort)(nil)

type selectionSort struct{}

func (s *selectionSort) sort(arr []int) {
	fmt.Println("called selection sort")
}
