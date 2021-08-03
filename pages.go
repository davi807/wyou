package main

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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

var resultContainer container.Scroll
var result *fyne.Container
var formatsContainer *fyne.Container

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

	result = container.NewVBox()

	resultContainer := container.NewScroll(result)
	resultContainer.SetMinSize(fyne.NewSize(APP_WIDTH-20, APP_HEIGHT-120))

	content = container.NewVBox(
		titleInfo,
		urlBar,
		searchButton,
		progress,
		resultContainer,
	)

	SetWindowContent(content)
}

func createVideoPage() {
	result.Objects = []fyne.CanvasObject{}

	//Create video informarion
	topBar := container.NewGridWithColumns(2)
	result.Add(topBar)

	// Add extractor specific llinks
	providerSpecials := createProviderSpecial()

	if providerSpecials != nil {
		// result.Add(providerSpecials)
	}

	// Add download formats
	formatsContainer = container.NewVBox()
	result.Add(formatsContainer)

	imageContainer := container.NewGridWrap(fyne.NewSize(THUMBNAIL_WIDTH, THUMBNAIL_HEIGHT))
	infoContainer := container.NewVBox()

	title := widget.NewLabel(video.Title)
	title.Alignment = fyne.TextAlignCenter
	title.Wrapping = fyne.TextWrapWord
	title.TextStyle.Bold = true
	infoContainer.Add(title)

	infoContainer.Add(widget.NewLabel("Duration: " + formatDuration(video.Duration)))
	infoContainer.Add(widget.NewLabel("Provider: " + video.Extractor))
	infoContainer.Add(createDirectDownload())

	loading := container.NewCenter(widget.NewLabel("Loading image..."))
	imageContainer.Add(loading)

	renderThumbnail := func() {
		resource, err := fyne.LoadResourceFromURLString(getVideoThumbnail(video).Url)
		if err != nil {
			return
		}

		img := canvas.NewImageFromResource(resource)
		img.SetMinSize(fyne.NewSize(THUMBNAIL_WIDTH-5, THUMBNAIL_HEIGHT-5))

		imageContainer.Remove(loading)

		imageContainer.Add(img)
		imageContainer.Refresh()
	}

	topBar.Add(imageContainer)
	topBar.Add(infoContainer)

	for _, links := range createDownloadabelItems(video.Formats) {
		formatsContainer.Add(links)
	}

	go time.AfterFunc(1*time.Second, renderThumbnail)
	searchButton.Enable()
	result.Show()
	// resultContainer.Refresh()
}

// Helper functions //

/* Get video image*/

func getVideoThumbnail(vd videoData) thumbnail {

	exts := []string{"jpg", "png", "bmp", "gif"}

	thumbMainUrl := thumbnail{Url: vd.ThumbnailURL}

	thumbs := append([]thumbnail{thumbMainUrl}, vd.Thumbnails...)

	sort.Slice(
		thumbs,
		func(i, j int) bool { return thumbs[j].Width < thumbs[i].Width },
	)

	for _, thumb := range thumbs {
		thurl, err := url.Parse(thumb.Url)

		if err != nil {
			continue
		}

		imageExt := strings.Split(thurl.Path, ".")
		if len(imageExt) == 0 {
			continue
		}

		ext := strings.ToLower(imageExt[len(imageExt)-1])

		for _, ex := range exts {
			if ext == ex {
				if video.Extractor == "Youtube" {
					thumb.Url = strings.Replace(thumb.Url, thurl.RawQuery, "", 1)
				}
				return thumb
			}
		}

	}

	return thumbnail{Url: THUMBNAIL_DEFAULT_URL}
}

/* Content set on window */

func SetWindowContent(cnt *fyne.Container) {
	window.SetContent(
		container.NewGridWrap(fyne.NewSize(APP_WIDTH, APP_HEIGHT), cnt),
	)
}

/* format video duration */

func formatDuration(d int) string {

	res := ""
	spr := ":"

	hours := d / 3600
	rems := d % 3600
	minutes := rems / 60
	seconds := d % 60

	if hours > 0 {
		res += fmt.Sprintf("%d", hours) + spr
	}

	res += fmt.Sprintf("%02d", minutes) + spr
	res += fmt.Sprintf("%02d", seconds)

	return res
}
