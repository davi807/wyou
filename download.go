package main

import (
	"os"
	"os/exec"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var downloadDir string
var videoFileName string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "c:\\"
	}
	downloadDir = path.Join(homeDir, "Downloads")
	videoFileName = path.Join(downloadDir, OUTPUT_FILE_NAME)

	stat, err := os.Stat(downloadDir)
	if err != nil || !stat.IsDir() {
		downloadDir = homeDir
	}
}

func autoDownload() {

	onDownloadClick(format{Ext: video.Ext})
}

func onDownloadClick(f format) {
	searchButton.Disable()
	progress.Show()
	result.Hide()

	arguments := []string{
		"--load-info-json", currentJsonFile,
		"--output", videoFileName,
	}

	if f.Id != "" {
		arguments = append(arguments, "--format", f.Id)
	}

	downloadProcess := exec.Command(YT_NAME, arguments...)

	downloadProcess.Run()

	resText := widget.NewLabel("Check file:\n" + path.Join(downloadDir, video.Title+"."+f.Ext))
	resContent := container.NewGridWrap(fyne.Size{Width: APP_WIDTH * 0.8, Height: APP_HEIGHT * 0.4}, resText)

	dialog.ShowCustom("Download finished", "Close", resContent, window)

	result.Show()
	progress.Hide()
	searchButton.Enable()
}
