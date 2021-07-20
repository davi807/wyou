package main

import (
	"errors"
	"fmt"
	"os/exec"

	"fyne.io/fyne/v2/dialog"
)

var jsonData = []byte{}

func onInfoClick() {
	progress.Show()
	cmd := exec.Command(YT_NAME, "-J", "https://www.youtube.com/watch?v=qHm9MG9xw1o")

	output, err := cmd.CombinedOutput()
	if err != nil {
		dialog.ShowError(err, window)
		return
	}

	if cmd.ProcessState.ExitCode() != 0 {
		dialog.ShowError(errors.New(string(output[0:256])), window)
		return
	}

	fmt.Println(len(output))
	/*
		f := formatBase{}

		json.Unmarshal(output, &f)
		fmt.Println(f)

	*/
}
