package main

import (
	"path/filepath"
	"runtime"
)

var YT_NAME = "youtube-dl"
var FF_NAME = "ffmpeg"

func init() {
	if runtime.GOOS == "windows" {
		YT_NAME = filepath.Join("bin", YT_NAME)
		FF_NAME = filepath.Join("bin", FF_NAME)
	}
}

const THUMBNAIL_WIDTH = 336
const THUMBNAIL_HEIGHT = 188

const THUMBNAIL_DEFAULT_URL = "https://via.placeholder.com/336x188.png?text=Image+not+found"

const OUTPUT_FILE_NAME = "%(title)s.%(ext)s"
