package main

import (
	"net"
	"os"
)

func testDNS() {
	names := []string{
		"roughtime.cloudflare.com",
		"cl.cam.ac.uk",
	}
	for _, name := range names {
		wg.Add(1)
		go checkDNS(name)
	}
	wg.Wait()
}

func checkDNS(name string) {
	_, err := net.LookupTXT(name)
	if err != nil {
		os.Exit(3)
	}
	wg.Done()
}
