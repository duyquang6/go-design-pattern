package main

type adapter struct {
	legacy ILegacy
}

func (a *adapter) callDB() bool {
	return a.legacy.callDB() != 0
}
