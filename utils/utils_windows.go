//go:build windows
// +build windows

package utils

import (
	"errors"

	"golang.org/x/sys/windows"
)

func IsUnreachableError(err error) bool {
	return errors.Is(err, windows.WSAENETUNREACH)
}
