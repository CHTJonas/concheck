package main

import (
	"net/http"
	"os"
)

var client = &http.Client{}

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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cache-Control", "no-store, max-age=0")
	req.Header.Set("User-Agent", "concheck/"+version+" (+https://github.com/CHTJonas/concheck)")
	resp, err := client.Do(req)
	if err != nil {
		os.Exit(1)
	}
	resp.Body.Close()
	wg.Done()
}
