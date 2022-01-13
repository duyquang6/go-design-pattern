package main

import "fmt"

var _ unstableSort = (*heapsort)(nil)

type heapsort struct{}

func (s *heapsort) sort(arr []int) {
	fmt.Println("called heap sort")
}
