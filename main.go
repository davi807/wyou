package main

import "net/http"

func main() {
	prepareServer()
	println("Server working on\n" + "http://" + address + ":" + port)
	http.ListenAndServe(address+":"+port, nil)
}
