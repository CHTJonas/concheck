package main

import (
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func check(url string) {
	_, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	wg.Done()
}

func main() {
	URLs := []string{"https://captive.apple.com/", "https://connectivitycheck.gstatic.com/generate_204"}
	for _, url := range URLs {
		wg.Add(1)
		go check(url)
	}
	wg.Wait()
}
