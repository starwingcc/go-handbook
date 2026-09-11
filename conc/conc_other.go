//go:build !linux

package conc

func platformName() string { return "non-linux" }
