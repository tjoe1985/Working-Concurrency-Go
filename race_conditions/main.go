package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func main() {
	msg = "hello joel"
	wg.Add(2)
	go updateMessage("hello john")
	go updateMessage("hello johnathan")
	wg.Wait()
	fmt.Println(msg)
}
func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

/*
Chances are that we won't know we have a race condition on this program
unless we run  go run -race . or some other tool.
*/
