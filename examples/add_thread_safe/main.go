package main

import (
	"log"
	"strconv"
	"sync"

	"github.com/fatih/set"
)

func main() {
	log.Print("Thread safe set operations")

	log.Print("Define wait group for waiting on goroutines")
	var wg sync.WaitGroup

	log.Print("Initialize our thread safe Set")
	s := set.New(set.ThreadSafe)

	log.Print("Add items concurrently (item1, item2, and so on)")
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			item := "item" + strconv.Itoa(i)
			log.Print("adding " + item)
			s.Add(item)
		}(i)
	}

	log.Print("Wait until all concurrent calls finished and print our set")
	wg.Wait()
	log.Print(s)

	log.Print("Done")
}
