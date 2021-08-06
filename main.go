package main

import "github.com/webview/webview"

var window webview.WebView

func main() {
	debug := true

	window := webview.New(debug)
	defer window.Destroy()

	window.Bind("load", loader)
	window.SetTitle(APP_TITLE)
	window.SetSize(APP_WIDTH, APP_HEIGHT, webview.HintFixed)

	window.Navigate(initHTML())

	window.Run()
}
