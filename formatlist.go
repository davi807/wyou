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
		// get best quality
		vsize := float64(video.RequestedFormats[0].Size) / 1024 / 1024
		asize := float64(video.RequestedFormats[1].Size) / 1024 / 1024
		ext := video.RequestedFormats[0].Ext + "+" + video.RequestedFormats[1].Ext

		downloadBtn := widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {})

		best := widget.NewCard("Best quality", fmt.Sprintf("%.02fMB + %.02fMB,  %s", vsize, asize, ext), downloadBtn)

		// download audio and save as mp3

		downloadMp3 := widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {})
		mp3 := widget.NewCard("Save as mp3", "Download audio and convert to mp3", downloadMp3)

		// append buttons

		row := container.NewGridWithColumns(2, best, mp3)
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
