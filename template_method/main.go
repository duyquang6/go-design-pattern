package main

func main() {
	concreteObjectA := concreteA{}
	abstract := abstractTemplate{concreteObjectA}

	abstract.templateMethod()

	concreteObjectB := concreteB{}
	abstract = abstractTemplate{concreteObjectB}
	abstract.templateMethod()
}
