package main

import (
	"net/http"
	"os"
)

func main() {
	URLs := []string{"https://captive.apple.com/", "https://connectivitycheck.gstatic.com/generate_204"}
	for _, url := range URLs {
		_, err := http.Get(url)
		if err != nil {
			os.Exit(1)
		}
	}
}
