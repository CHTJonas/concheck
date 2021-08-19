package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go testICMP()

	wg.Add(1)
	go testHTTP()

	wg.Wait()
}
