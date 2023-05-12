package main

import "fmt"

func main() {
	go PrintSomething("Print this first!")

	PrintSomething("Print this second!")
}
func PrintSomething(s string) {
	fmt.Println(s)
}
