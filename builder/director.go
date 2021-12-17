package main

type director struct {
	builder Builder
}

func (d *director) changeBuilder(builder Builder) {
	d.builder = builder
}

func (d *director) make(builderType string) object {
	if builderType == "A" {
		d.builder.buildPart1()
		d.builder.buildPart2()
	} else {
		d.builder.buildPart1()
		d.builder.buildPart2()
		d.builder.buildPart3()
	}
	return d.builder.getResult()
}
