package main

import "fmt"

func main() {
	legacy := legacy{}
	var client IClient = &adapter{&legacy}
	fmt.Println("adapter result:", client.callDB())
}
