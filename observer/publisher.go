package main

type concretePublisher struct {
	subscribers []subscriber
}

type msg struct{ msg string }

type publisher interface {
	subscribe(subscriber subscriber)
	unsubscribe(subscriber subscriber)
	notifySubscribers(msg msg)
}

func (p *concretePublisher) subscribe(subscriber subscriber) {
	// if subscriber is subcribbed
	for _, sub := range p.subscribers {
		if sub == subscriber {
			return
		}
	}

	p.subscribers = append(p.subscribers, subscriber)
}

func (p *concretePublisher) unsubscribe(subscriber subscriber) {
	for i, sub := range p.subscribers {
		if sub == subscriber {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
		}
	}
}

func (p *concretePublisher) notifySubscribers(msg msg) {
	for _, sub := range p.subscribers {
		sub.update(msg)
	}
}
