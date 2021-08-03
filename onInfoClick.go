package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var currentJsonFile string
var video videoData

func onInfoClick() {
	searchButton.Disable()

	result.Hide()
	progress.Show()
	defer progress.Hide()

	cmd := exec.Command(YT_NAME, "-J", "--no-playlist", urlBar.Text)

	jsonOutput, errCode := cmd.CombinedOutput()

	if errCode != nil {

		errorText := widget.NewLabel(string(jsonOutput))
		errorText.Wrapping = fyne.TextWrapWord

		errContent := container.NewGridWrap(fyne.Size{Width: APP_WIDTH * 0.8, Height: APP_HEIGHT * 0.4}, errorText)

		dialog.ShowCustom("Error: "+errCode.Error(), "Close", errContent, window)
		searchButton.Enable()
		return
	}

	video = videoData{}
	json.Unmarshal(jsonOutput, &video)

	currentJsonFile = path.Join(os.TempDir(), "wyou-current-video-info.tmp")
	tmpF, _ := os.Create(currentJsonFile)
	tmpF.Write(jsonOutput)

	go createVideoPage()
}
