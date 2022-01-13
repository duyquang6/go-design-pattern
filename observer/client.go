package main

func main() {
	var publisher publisher = &concretePublisher{}
	sub1 := &concreteSubscriber{index: 1}
	sub2 := &concreteSubscriber{index: 2}
	sub3 := &concreteSubscriber{index: 3}
	publisher.subscribe(sub1)
	publisher.subscribe(sub2)
	publisher.subscribe(sub3)

	publisher.notifySubscribers(msg{msg: "hello"})
}
