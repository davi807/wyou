package main

type formatBase struct {
	Extractor string `json:"extractor_key"`
	Title     string `json:"title"`
	Thumbnail string `json:"Thumbnail"`
	Duration  int    `json:"duration"`
}
