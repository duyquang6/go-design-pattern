package main

type button struct{ command Command }

func (b *button) onClick() {
	b.command.execute()
}
