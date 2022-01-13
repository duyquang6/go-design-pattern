package main

type op func(l *longStructDefinition)

type longStructDefinition struct {
	requiredA int
	requiredB string
	requiredC bool

	optA bool
	optB string
	optC int
}

func NewLongStructDefinition(requiredA int, requiredB string, requiredC bool, ops ...op) longStructDefinition {
	longStructDefinition := longStructDefinition{requiredA: requiredA, requiredB: requiredB, requiredC: requiredC}
	for _, opFnc := range ops {
		opFnc(&longStructDefinition)
	}
	return longStructDefinition
}

func WithOptA(optA bool) op {
	return func(l *longStructDefinition) {
		l.optA = optA
	}
}

func WithOptB(optB string) op {
	return func(l *longStructDefinition) {
		l.optB = optB
	}
}

func WithOptC(optC int) op {
	return func(l *longStructDefinition) {
		l.optC = optC
	}
}
