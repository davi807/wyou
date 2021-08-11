package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"
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

func download(format string) {
	arguments := []string{
		"--load-info-json", currentJsonFile,
		"--output", videoFileName,
	}

	if format != "best" {
		arguments = append(arguments, "--format", format)
	}

	downloadProcess := exec.Command(YT_NAME, arguments...)

	stdOut, err := downloadProcess.StdoutPipe()

	if err != nil {
		return
	}

	downloadProcess.Start()

	for downloadProcess.ProcessState.ExitCode() == -1 {
		buffer := make([]byte, 256)

		n, err := stdOut.Read(buffer)

		if err != nil {
			break
		}

		// fmt.Println()
		if n > 0 {
			stdOutChannel <- buffer[:n]
		}

		time.Sleep(1 * time.Second)
	}
	downloadProcess.Wait()
	println("end")
	close(stdOutChannel)

}
