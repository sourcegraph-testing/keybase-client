//go:build !windows
// +build !windows

package avatars

func fileUrlize(path string) string {
	return path
}
