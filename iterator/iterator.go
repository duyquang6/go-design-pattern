package main

import "errors"

type intIter struct {
	array []int
	curr  int
}

type IntIter interface {
	hasNext() bool
	next() (int, error)
}

func (_this *intIter) hasNext() bool {
	return _this.curr+1 < len(_this.array)
}

func (_this *intIter) next() (int, error) {
	if !_this.hasNext() {
		return 0, errors.New("exceed limit number of records")
	}
	_this.curr++
	return _this.array[_this.curr], nil
}
