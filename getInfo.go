package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
)

var currentJsonFile string

func getInfo(url string) []byte {

	cmd := exec.Command(YT_NAME, "-J", "--no-playlist", url)

	jsonOutput, errCode := cmd.CombinedOutput()

	if errCode != nil {

		errorText := string(jsonOutput)

		errData, _ := json.Marshal(map[string]interface{}{
			"error":     true,
			"errorText": errorText,
		})

		return errData
	}

	currentJsonFile = filepath.Join(os.TempDir(), "wyou-current-video-info.tmp")
	tmpF, _ := os.Create(currentJsonFile)
	tmpF.Write(jsonOutput)

	return jsonOutput
}
