package main

import (
	"log"
	"strconv"

	"github.com/fatih/set"
)

func main() {
	log.Print("Initialize our non thread safe Set")
	s := set.New(set.NonThreadSafe)

	log.Print("Add items serially (item1, item2, and so on)")
	for i := 0; i < 10; i++ {
		item := "item" + strconv.Itoa(i)
		log.Print("adding " + item)
		s.Add(item)
	}
	log.Print(s)

	log.Print("Done")
}
