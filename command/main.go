package main

import "fmt"

func main() {
	obj := object{}
	fmt.Println("click OK")
	fmt.Println("obj before click:", obj)
	okBtn := button{&okCommand{&obj}}
	okBtn.onClick()
	fmt.Println("obj after click:", obj)

	fmt.Println("click cancel")
	fmt.Println("obj before click:", obj)
	cancelBtn := button{&cancelCommand{&obj}}
	cancelBtn.onClick()
	fmt.Println("obj after click:", obj)
}
