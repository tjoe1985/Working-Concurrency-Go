package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Sup yo"
	wg.Add(2)

	go updateMessage("See you")
	go updateMessage("See ya")
	wg.Wait()
	if msg != "See ya" {
		t.Errorf("Incorect value in msg")
	}
}

/*
Chances are that we won't know we have a race condition on this program
unless we run  "go test -race ." or some other tool.
*/
