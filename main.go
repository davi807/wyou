package main

import "net/http"

func main() {
	prepareServer()
	makeServerHnadlers()

	println("Server working on\n" + "http://" + address + ":" + port)
	http.ListenAndServe(address+":"+port, nil)
}
