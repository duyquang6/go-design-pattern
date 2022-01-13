package main

var _ sortStrategy = (*intSort)(nil)

type intSort struct {
	sortStrategy unstableSort
}

type sortStrategy interface {
	setStrategy(strategy unstableSort)
	sort(arr []int)
}

func (s *intSort) setStrategy(strategy unstableSort) {
	s.sortStrategy = strategy
}

func (s *intSort) sort(arr []int) {
	s.sortStrategy.sort(arr)
}
