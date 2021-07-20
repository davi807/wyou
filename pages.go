package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var base fyne.App
var window fyne.Window
var content *fyne.Container

var titleInfo *widget.Label
var urlBar *widget.Entry
var searchButton *widget.Button

var progress fyne.Widget
var resultContainer fyne.Widget

func SetWindowContent(cnt *fyne.Container) {
	window.SetContent(
		container.NewGridWrap(fyne.Size{Width: APP_WIDTH, Height: APP_HEIGHT}, cnt),
	)
}

/* Pages */

func initStartPage() {
	base = app.New()
	window = base.NewWindow(APP_TITLE)

	titleInfo = widget.NewLabel("Enter video URL")
	// titleInfo.Wrapping = fyne.TextWrapWord

	urlBar = widget.NewEntry()
	urlBar.PlaceHolder = URLBAR_PLACEHOLDER

	searchButton = widget.NewButton("Get Info", onInfoClick)

	progress = widget.NewProgressBarInfinite()
	progress.Hide()

	content = container.NewVBox(
		titleInfo,
		urlBar,
		searchButton,
		progress,
	)

	SetWindowContent(content)
}
