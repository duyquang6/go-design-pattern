package main

type unstableSort interface {
	sort(arr []int)
}

type stableSort interface {
	sort(arr []int)
}
