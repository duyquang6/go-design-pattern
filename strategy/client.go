package main

func main() {
	heapSortStrategy := &heapsort{}
	quickSortStrategy := &quicksort{}
	selectionSortStrategy := &selectionSort{}

	var context sortStrategy = &intSort{}
	context.setStrategy(heapSortStrategy)

	context.sort([]int{1, 2, 3})

	context.setStrategy(quickSortStrategy)
	context.sort([]int{1, 2, 3})

	context.setStrategy(selectionSortStrategy)
	context.sort([]int{1, 2, 3})
}
