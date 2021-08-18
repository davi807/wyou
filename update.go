package main

import (
	"os/exec"
)

func update() []byte {

	updateProcess := exec.Command(YT_NAME, "-U")
	out, err := updateProcess.CombinedOutput()

	if err != nil {
		return []byte(err.Error())
	}

	return out
}
