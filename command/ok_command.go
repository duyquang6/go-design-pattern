package main

type okCommand struct{ object *object }

func (c *okCommand) execute() {
	c.object.value = true
}
