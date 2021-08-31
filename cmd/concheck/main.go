package main

import (
	"sync"
)

var (
	version = "dev-edge"
	wg      sync.WaitGroup
)

func main() {
	testICMP()
	testDNS()
	testHTTP()
}
