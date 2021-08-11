package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
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

func makeServerHnadlers() {
	http.HandleFunc("/api/get-info", func(rw http.ResponseWriter, r *http.Request) {

		body := make([]byte, 256)
		n, _ := r.Body.Read(body)
		url := string(body[:n])

		jsonResult := getInfo(string(url))

		rw.Write([]byte(jsonResult))
	})

	http.HandleFunc("/api/download/", func(rw http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.String(), "/")

		println(path[len(path)-1])
	})
}
