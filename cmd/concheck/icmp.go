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
		"212.187.216.254", "2001:1900:5:2:2:0:11c:2f2", // JANET.ear1.London1.Level3.net
		"46.227.201.1", "2a01:9e00::201:1", // lo.aebi.m.faelix.net
		"193.0.0.164", "2001:67c:2e8:3::c100:a4", // ping.ripe.net
	}
	unreachableCount := 0
	for _, endpoint := range endpoints {
		if (forceIPv4Flag && !utils.IsIPv4(endpoint)) || (forceIPv6Flag && utils.IsIPv4(endpoint)) {
			continue
		}
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
			// If the user has explicitly forced IPv4 or IPv6 then don't squelch errors
			if forceIPv4Flag || forceIPv6Flag {
				os.Exit(2)
			}
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
