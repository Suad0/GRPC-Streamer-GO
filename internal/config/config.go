package config

import (
	"os"
)

func GetVideoDir() string {
	videoDir := os.Getenv("VIDEO_DIR")
	if videoDir == "" {
		return "./video/"
	}
	return videoDir
}
