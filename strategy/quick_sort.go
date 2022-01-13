package main

import "fmt"

var _ unstableSort = (*quicksort)(nil)

type quicksort struct{}

func (s *quicksort) sort(arr []int) {
	fmt.Println("called quick sort")
}
