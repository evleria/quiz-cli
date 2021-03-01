package config

import (
	"os"
	"strings"
)

func ExpandPath(path string) string {
	if !strings.HasPrefix(path, "~") {
		return path
	}

	return os.Getenv("HOME") + path[1:]
}
