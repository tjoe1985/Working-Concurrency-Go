package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_PrintSomething(t *testing.T) {
	stdOut := os.Stdout

	read, write, _ := os.Pipe()
	//chage the os standar output to my write
	os.Stdout = write
	//create waitgroup for my test
	var wg sync.WaitGroup
	wg.Add(1)

	//call function to test
	go PrintSomething("hey", &wg)
	//wait for the function to be done
	wg.Wait()

	// close our pipe ignoring the error
	_ = write.Close()
	// get the result from the function in var result
	result, _ := io.ReadAll(read)
	// change in to readable string by casting it.
	output := string(result)
	// reset overwriten vars

	os.Stdout = stdOut
	/*
				Perform the actual test
			if the string does not contain the passed string return the error
		    else we should be good.
	*/
	if !strings.Contains(output, "hey") {
		t.Errorf("Expected to find hey , but did not. ")
	}
}
