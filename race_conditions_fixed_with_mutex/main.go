package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func main() {
	msg = "hello joel"
	var mutex sync.Mutex
	wg.Add(2)
	go updateMessage("hello john", &mutex)
	go updateMessage("hello johnathan", &mutex)
	wg.Wait()
	fmt.Println(msg)
}
func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	// lock and unloc the resource to about race condition
	m.Lock()
	msg = s
	m.Unlock()
}

/*
Chances are that we won't know we have a race condition on this program
unless we run  go run -race . or some other tool.
*/
