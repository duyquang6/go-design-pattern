package main

type AbstractTemplate interface {
	step1()
	step2()
	step3()
}

type abstractTemplate struct {
	AbstractTemplate
}

func (c abstractTemplate) templateMethod() {
	c.step1()
	c.step2()
	c.step3()
}
