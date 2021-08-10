package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

var address = "localhost"
var port string

func prepareServer() {
	// getting any free port
	conn, err := net.Listen("tcp", address+":0")
	if err != nil {
		log.Fatal(err)
	}

	port = fmt.Sprint(conn.Addr().(*net.TCPAddr).Port)
	conn.Close()

	// making static file server on root location
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/", fs)
}
