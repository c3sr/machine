package osinfo

import "sync"

var (
	info string
)

func Info() string {
	var once sync.Once
	once.Do(func() {
		info = localOSVersion()
	})
	return info
}
