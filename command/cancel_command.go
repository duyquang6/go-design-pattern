package main

type cancelCommand struct{ object *object }

func (c *cancelCommand) execute() {
	c.object.value = false
}
