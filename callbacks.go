package main

import (
	"encoding/json"
	"net/url"
	"sort"
	"strings"

	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var jsonData = []byte{}

func onInfoClick() {
	/*
		go func() {
			resp, _ := http.Get("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ-Pia31MHkdochSVtOiP2P_k8rjBPtGl0Uvg&usqp=CAU")
			img := canvas.NewImageFromReader(resp.Body, "afadas")

			img.SetMinSize(fyne.Size{320, 240})
			icont := container.NewGridWrap(fyne.Size{320, 240}, img)

			content.Add(icont)
		}()
	*/

	progress.Show()
	defer progress.Hide()

	cmd := exec.Command(YT_NAME, "-J", "--no-playlist", urlBar.Text)

	output, errCode := cmd.CombinedOutput()

	if errCode != nil {

		errorText := widget.NewLabel(string(output))
		errorText.Wrapping = fyne.TextWrapWord

		errContent := container.NewGridWrap(fyne.Size{Width: APP_WIDTH * 0.8, Height: APP_HEIGHT * 0.4}, errorText)

		dialog.ShowCustom("Error: "+errCode.Error(), "Close", errContent, window)
		return
	}

	video := videoData{}
	json.Unmarshal(output, &video)
}

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
				return thumb
			}
		}

	}

	return thumbnail{}
}
