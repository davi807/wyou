package main

type videoData struct {
	Id               string      `json:"id"`
	Extractor        string      `json:"extractor_key"`
	Title            string      `json:"title"`
	ThumbnailURL     string      `json:"thumbnail"`
	Thumbnails       []thumbnail `json:"thumbnails"`
	Duration         int         `json:"duration"`
	Ext              string      `json:"ext"`
	BestFormatId     string      `json:"format_id"`
	Formats          []format    `json:"formats"`
	RequestedFormats []format    `json:"requested_formats"`
}

type format struct {
	Size       uint64 `json:"filesize"`
	Id         string `json:"format_id"`
	Resolution string `json:"resolution"`

	// Size in bytes
	Format string `json:"format"`

	Acodec string `json:"acodec"`
	Vcodec string `json:"vcodec"`
	Ext    string `json:"ext"`
}

type thumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
