package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("WYou")

	hello := widget.NewLabel("Video downloader")
	button := widget.NewButton("Hi!", func() {
		hello.SetText("Welcome :)")
	})

	w.SetContent(container.NewGridWrap(fyne.Size{Width: 680, Height: 400},
		container.NewVBox(
			hello,
			button,
		)),
	)

	w.ShowAndRun()
}
