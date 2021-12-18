package main

type collection interface {
	createIterator() IntIter
}

type intCollection struct {
	arr []int
}

func (i *intCollection) createIterator() IntIter {
	return &intIter{i.arr, -1}
}
