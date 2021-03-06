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
var downloading bool

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

		var stdOutChannel chan []byte

		downloading = true
		defer func() {
			downloading = false
		}()

		path := strings.Split(r.URL.String(), "/")
		format := path[len(path)-1]
		stdOutChannel = make(chan []byte)
		go download(format, stdOutChannel)

		flusher := rw.(http.Flusher)

		for row := range stdOutChannel {
			rw.Write(row)
			flusher.Flush()
		}
		rw.Write([]byte("##DONE##"))
	})

	http.HandleFunc("/api/update/", func(rw http.ResponseWriter, r *http.Request) {

		res := update()

		rw.Write(res)
		rw.Write([]byte("##DONE##"))
	})
}
