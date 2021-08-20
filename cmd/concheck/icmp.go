package main

import (
	"os"
	"time"

	"github.com/go-ping/ping"
)

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
