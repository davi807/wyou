package main

type videoData struct {
	Extractor    string   `json:"extractor_key"`
	Title        string   `json:"title"`
	Thumbnail    string   `json:"thumbnail"`
	Duration     int      `json:"duration"`
	BestFormatId string   `json:"format_id"`
	Formats      []format `json:"formats"`
}

type format struct {
	Size     uint64 `json:"filesize"`
	FormatId string `json:"format_id"`
}
