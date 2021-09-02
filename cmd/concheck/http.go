package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"os"
	"time"
)

func testHTTP() {
	URLs := []string{
		"https://chtj2.user.srcf.net/static",
		"https://connectivitycheck.gstatic.com/generate_204",
		"https://cloudflare.com/cdn-cgi/trace",
	}
	URLs = append(flag.Args(), URLs...)
	for _, url := range URLs {
		wg.Add(1)
		go checkHTTP(url)
	}
	wg.Wait()
}

func checkHTTP(url string) {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
				if forceIPv4Flag {
					network += "4"
				} else if forceIPv6Flag {
					network += "6"
				}
				return (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext(ctx, network, address)
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
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
