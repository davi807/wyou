package main

type videoData struct {
	Id           string      `json:"id"`
	Extractor    string      `json:"extractor_key"`
	Title        string      `json:"title"`
	ThumbnailURL string      `json:"thumbnail"`
	Thumbnails   []thumbnail `json:"thumbnails"`
	Duration     int         `json:"duration"`
	Ext          string      `json:"ext"`
	BestFormatId string      `json:"format_id"`
	Formats      []format    `json:"formats"`
}

type format struct {
	Size     uint64 `json:"filesize"`
	FormatId string `json:"format_id"`
}

type thumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
