package main

type IClient interface {
	callDB() bool
}
