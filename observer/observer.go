package main

import "fmt"

type concreteSubscriber struct{ index int }

type subscriber interface {
	update(msg msg)
}

func (s *concreteSubscriber) update(msg msg) {
	fmt.Println("Observer", s.index, "update with message:", msg)
}
