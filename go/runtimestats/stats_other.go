//go:build !darwin
// +build !darwin

package runtimestats

func getStats() (res statsResult) {
	return res
}
