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

// i am repeating myself :((
func autoDownload() {
	searchButton.Disable()
	progress.Show()
	result.Hide()

	downloadProcess := exec.Command(YT_NAME,
		"--load-info-json", currentJsonFile,
		"--output", videoFileName,
	)

	out, _ := downloadProcess.CombinedOutput()

	resText := widget.NewLabel(string(out))
	resContent := container.NewGridWrap(fyne.Size{Width: APP_WIDTH * 0.8, Height: APP_HEIGHT * 0.4}, resText)

	dialog.ShowCustom("Download finished", "Close", resContent, window)

	result.Show()
	progress.Hide()
	searchButton.Enable()

}

func onDownladClick(f format) {
	searchButton.Disable()
	progress.Show()
	result.Hide()

	downloadProcess := exec.Command(YT_NAME,
		"--load-info-json", currentJsonFile,
		"--format", f.Id,
		"--output", videoFileName,
	)

	out, _ := downloadProcess.CombinedOutput()

	resText := widget.NewLabel(string(out))
	resContent := container.NewGridWrap(fyne.Size{Width: APP_WIDTH * 0.8, Height: APP_HEIGHT * 0.4}, resText)

	dialog.ShowCustom("Download finished", "Close", resContent, window)

	result.Show()
	progress.Hide()
	searchButton.Enable()
}
