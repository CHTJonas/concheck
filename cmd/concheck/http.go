package main

import (
	"net/http"
	"os"
)

func testHTTP() {
	URLs := []string{
		"https://chtj2.user.srcf.net/static",
		"https://connectivitycheck.gstatic.com/generate_204",
		"https://cloudflare.com/cdn-cgi/trace",
	}
	URLs = append(os.Args[1:], URLs...)
	for _, url := range URLs {
		wg.Add(1)
		go checkHTTP(url)
	}
	wg.Wait()
}

func checkHTTP(url string) {
	_, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	wg.Done()
}
