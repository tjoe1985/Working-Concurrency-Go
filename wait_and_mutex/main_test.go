package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	// get standar input output
	stdOut := os.Stdout
	read, write, _ := os.Pipe()
	//chage the os standar output to my write
	os.Stdout = write
	main()
	// close our pipe ignoring the error
	_ = write.Close()
	// get the result from the function in var result
	result, _ := io.ReadAll(read)
	// change in to readable string by casting it.
	output := string(result)
	// reset overwriten vars

	os.Stdout = stdOut
	/*
		we know the value of our final balance should be 34320 for now cause we know
		how much our hard coded values are so for now we can test for that.

	*/
	if !strings.Contains(output, "$34320.00") {
		t.Error("Expected to find balance of 34320  , but did not. ")
	}
}

//go test -race -cover -v .
