package main

import (
	"fmt"
	"runtime"
	"strings"
)

func uaString() string {
	return fmt.Sprintf("Concheck/%s Go/%s (+https://github.com/CHTJonas/concheck)",
		version, strings.TrimPrefix(runtime.Version(), "go"))
}
