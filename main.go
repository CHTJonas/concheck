package main

import (
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

var wg sync.WaitGroup

func checkHTTP(url string) {
	_, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	wg.Done()
}

func testHTTP() {
	URLs := []string{
		"https://chtj2.user.srcf.net/static",
		"https://captive.apple.com/",
		"https://connectivitycheck.gstatic.com/generate_204",
		"https://cloudflare.com/cdn-cgi/trace",
	}
	for _, url := range URLs {
		wg.Add(1)
		go checkHTTP(url)
	}
	wg.Done()
}

func checkICMP(endpoint string) {
	pinger, err := ping.NewPinger(endpoint)
	if err != nil {
		panic(err)
	}
	pinger.Count = 5
	pinger.Interval = 100 * time.Millisecond
	pinger.Timeout = 50 * time.Second
	// pinger.SetPrivileged(true)
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	if stats.PacketLoss > 0.5 {
		os.Exit(2)
	}
	wg.Done()
}

func testICMP() {
	endpoints := []string{
		"www.jump.net.uk",
		"one.one.one.one",
		"dns.google",
		"dns.quad9.net",
	}
	for _, endpoint := range endpoints {
		wg.Add(1)
		go checkICMP(endpoint)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go testHTTP()

	wg.Add(1)
	go testICMP()

	wg.Wait()
}
