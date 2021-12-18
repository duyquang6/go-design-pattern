package main

type legacy struct{}

type ILegacy interface {
	callDB() int
}

func (l *legacy) callDB() int {
	return 1
}
