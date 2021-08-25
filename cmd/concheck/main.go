package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	testICMP()
	testDNS()
	testHTTP()
}
