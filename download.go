package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

var downloadDir string
var videoFileName string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "c:\\"
	}
	downloadDir = filepath.Join(homeDir, "Downloads")
	videoFileName = filepath.Join(downloadDir, OUTPUT_FILE_NAME)

	stat, err := os.Stat(downloadDir)
	if err != nil || !stat.IsDir() {
		downloadDir = homeDir
	}
}

func download(format string, info ...string) string {
	title := info[0]
	extension := info[1]

	arguments := []string{
		"--load-info-json", currentJsonFile,
		"--output", videoFileName,
	}

	if format != "best" {
		arguments = append(arguments, "--format", format)
	}

	downloadProcess := exec.Command(YT_NAME, arguments...)

	downloadProcess.Run()

	return filepath.Join(downloadDir, title+"."+extension)
}
