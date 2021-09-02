package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

var (
	version       = "dev-edge"
	wg            sync.WaitGroup
	forceIPv4Flag bool
	forceIPv6Flag bool
)

func init() {
	flag.Usage = func() {
		fmt.Println(usage)
	}
	flag.BoolVar(&forceIPv4Flag, "4", false, "force the use of IPv4")
	flag.BoolVar(&forceIPv6Flag, "6", false, "force the use of IPv6")
	flag.Parse()
}

func main() {
	if forceIPv4Flag && forceIPv6Flag {
		fmt.Println("concheck: cannot specify both -4 and -6")
		os.Exit(125)
	}
	testICMP()
	testDNS()
	testHTTP()
}
