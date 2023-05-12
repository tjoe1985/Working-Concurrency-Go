package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	for _, s := range numbers {
		//add a routine to the wait group so it know how many to wait for.
		wg.Add(1)
		go PrintSomething(s, &wg)
	}
	//asking main with wait groups to wait for other routines to finish
	wg.Wait()
	//compensate for negative WaitGroup counter.
	wg.Add(1)
	PrintSomething("Print this second!", &wg)

	excTime := time.Since(start)
	log.Println("Time executing: ", excTime)
}
func PrintSomething(s string, wg *sync.WaitGroup) {
	//decrement wait group by one after code runs
	defer wg.Done()
	fmt.Println(s)
}
