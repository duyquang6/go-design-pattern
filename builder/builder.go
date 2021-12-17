package main

type Builder interface {
	buildPart1()
	buildPart2()
	buildPart3()
	getResult() object
}

type builderA struct {
	object
}

func (b *builderA) buildPart1() {
	b.a = 1
}

func (b *builderA) buildPart2() {
	b.b = "b"
}

func (b *builderA) buildPart3() {
	b.c = true
}

func (b *builderA) getResult() object {
	return object{b.a, b.b, b.c}
}

type builderB struct {
	object
}

func (b *builderB) buildPart1() {
	b.a = 2
}

func (b *builderB) buildPart2() {
	b.b = "c"
}

func (b *builderB) buildPart3() {
	b.c = false
}

func (b *builderB) getResult() object {
	return object{b.a, b.b, b.c}
}
