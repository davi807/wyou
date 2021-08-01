package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createProviderSpecial() *fyne.Container {
	res := container.NewVBox()

	switch video.Extractor {
	case "Youtube":

		// download audio and save as mp3

		downloadMp3 := widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {})
		mp3Title := widget.NewLabel("Save as mp3")
		mp3Title.TextStyle.Bold = true
		mp3Subtitle := widget.NewLabel("Download audio and convert to mp3")
		mp3 := container.NewVBox(mp3Title, mp3Subtitle, downloadMp3)

		// append buttons

		row := container.NewVBox(mp3)
		res.Add(row)

	default:
		return nil
	}

	return res
}

func createDownloadabelItems(formats []format) []*fyne.Container {
	res := []*fyne.Container{}

	// var addedInList bool

	var formatsRow *fyne.Container
	var hasUnpushed bool

	// if youtube change order for audio+video item set on top place on list
	if video.Extractor == "Youtube" {
		formats[0] = formats[len(formats)-1]
	}

	for index := 0; index < len(formats); index++ {
		formatItem := formats[index]

		hasUnpushed = true
		if index%2 == 0 {
			formatsRow = container.NewGridWithColumns(2)
		}

		videoItem := container.NewVBox()

		title := widget.NewLabel(formatItem.Ext + " / " + fmt.Sprintf("%.02fMB", float64(formatItem.Size)/1024/1024))
		title.TextStyle.Bold = true
		videoItem.Add(title)

		codecText := "Audio: " + formatItem.Acodec + " ; video: " + formatItem.Vcodec
		codecText += "\nFordmat: " + formatItem.Format

		videoItem.Add(widget.NewLabel(codecText))

		videoItem.Add(widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {
			fmt.Println(index)
		}))

		formatsRow.Add(videoItem)

		if index%2 == 1 {
			res = append(res, formatsRow)
			hasUnpushed = false
		}

	}

	if hasUnpushed {
		res = append(res, formatsRow)
	}

	return res
}

func createDirectDownload() fyne.CanvasObject {
	return widget.NewButtonWithIcon("Download best quality", theme.DownloadIcon(), func() {

	})
}
