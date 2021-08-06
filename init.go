package main

import (
	"io/ioutil"
	"log"
	"path"
)

func initHTML() string {
	html := "data:text/html,"

	indexURI := path.Join("assets", "loader.html")
	index, err := ioutil.ReadFile(indexURI)

	if err != nil {
		log.Fatal(err)
	}

	html += string(index)

	return html
}

func loader(name string) (string, error) {
	res, err := ioutil.ReadFile("./assets" + name)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
