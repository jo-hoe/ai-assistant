package cmd

import (
	"fmt"

	webview "github.com/webview/webview_go"
)

func StartWebview(port int) {
	view := webview.New(false)
	defer view.Destroy()
	view.SetTitle("AI Assistant")
	view.SetSize(1024, 600, webview.HintNone)
	view.Navigate(fmt.Sprintf("http://localhost:%d", port))
	view.Run()
}
