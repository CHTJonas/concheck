//go:build linux || freebsd || openbsd || darwin
// +build linux freebsd openbsd darwin

package utils

import (
	"errors"

	"golang.org/x/sys/unix"
)

func IsUnreachableError(err error) bool {
	return errors.Is(err, unix.ENETUNREACH) || errors.Is(err, unix.EHOSTUNREACH)
}
