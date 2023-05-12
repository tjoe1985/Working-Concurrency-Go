package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	read, write, _ := os.Pipe()
	//chage the os standar output to my write
	os.Stdout = write
	msg = "hey"
	//call function to test
	printMessage()
	//wait for the function to be done

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

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage("hey", &wg)
	wg.Wait()

	if msg != "hey" {
		t.Errorf("Expected to find hey , but did not.")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	read, write, _ := os.Pipe()
	//chage the os standar output to my write
	os.Stdout = write

	//call function to test
	main()
	//wait for the function to be done

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
	fmt.Println("value of output: ", output)
	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected to find Hello, universe! , but got: %s ", output)
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected to find Hello, cosmos! , but got: %s ", output)
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find Hello, world! , but but got: %s", output)
	}
}
