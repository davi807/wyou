package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var currentJsonFile string
var video videoData

func onInfoClick() {
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
		return
	}

	video = videoData{}
	json.Unmarshal(jsonOutput, &video)

	currentJsonFile = os.TempDir() + "/" + video.Id + time.Now().Format("150405.000") + ".tmp"
	tmpF, _ := os.Create(currentJsonFile)
	tmpF.Write(jsonOutput)

	go createVideoPage()
}
