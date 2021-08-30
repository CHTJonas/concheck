package main

import (
	"os"
	"runtime"
	"time"

	"github.com/CHTJonas/concheck/utils"
	"github.com/go-ping/ping"
)

func testICMP() {
	endpoints := []string{
		"1.0.0.1", "2606:4700:4700::1001",
		"8.8.4.4", "2001:4860:4860::8844",
		"149.112.112.112", "2620:fe::fe",
	}
	unreachableCount := 0
	for _, endpoint := range endpoints {
		wg.Add(1)
		go checkICMP(endpoint, &unreachableCount)
	}
	wg.Wait()
	if unreachableCount > 3 {
		os.Exit(2)
	}
}

func checkICMP(endpoint string, unreachableCount *int) {
	pinger, err := ping.NewPinger(endpoint)
	if err != nil {
		panic(err)
	}
	pinger.Count = 5
	pinger.Interval = 100 * time.Millisecond
	pinger.Timeout = 5 * time.Second
	if runtime.GOOS != "darwin" {
		pinger.SetPrivileged(true)
	}
	err = pinger.Run()
	if err != nil {
		if utils.IsUnreachableError(err) {
			*unreachableCount++
			wg.Done()
			return
		} else {
			panic(err)
		}
	}
	stats := pinger.Statistics()
	if stats.PacketLoss > 0.5 {
		os.Exit(2)
	}
	wg.Done()
}
