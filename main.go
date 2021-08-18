package main

import (
	"net/http"
	"os/exec"
	"time"
)

func main() {
	prepareServer()
	makeServerHnadlers()

	url := "http://" + address + ":" + port
	println("Server working on\n" + url)
	println("Opening browser...")

	go func() {
		time.Sleep(500 * time.Microsecond)
		exec.Command(DEFAULT_BROWSER, url).Output()
		println("\nPLEASE NOT CLOSE THIS PROGRAM")
	}()

	http.ListenAndServe(address+":"+port, nil)
}
