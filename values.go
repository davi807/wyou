package main

import (
	"path/filepath"
	"runtime"
)

var YT_NAME = "youtube-dl"
var FF_NAME = "ffmpeg"

var DEFAULT_BROWSER = "xdg-open"

func init() {
	if runtime.GOOS == "windows" {
		YT_NAME = filepath.Join("bin", YT_NAME)
		FF_NAME = filepath.Join("bin", FF_NAME)
		DEFAULT_BROWSER = "explorer"
	}
}

const OUTPUT_FILE_NAME = "%(title)s.%(ext)s"
