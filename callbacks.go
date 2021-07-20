package main

import (
	"encoding/json"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var jsonData = []byte{}

func onInfoClick() {
	progress.Show()
	defer progress.Hide()

	cmd := exec.Command(YT_NAME, "-J", "--no-playlist", urlBar.Text)

	output, errCode := cmd.CombinedOutput()

	if errCode != nil {

		errorText := widget.NewLabel(string(output))
		errorText.Wrapping = fyne.TextWrapWord

		errContent := container.NewGridWrap(fyne.Size{APP_WIDTH * 0.8, APP_HEIGHT * 0.4}, errorText)

		dialog.ShowCustom("Error: "+errCode.Error(), "Close", errContent, window)
		return
	}

	f := videoData{}
	json.Unmarshal(output, &f)

	name := widget.NewLabel(f.Title)
	content.Add(name)
}
