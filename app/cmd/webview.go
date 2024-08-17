package cmd

import (
	"fmt"

	webview "github.com/webview/webview_go"
)

func StartWebview(port int) {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("AI Assistant")
	w.SetSize(1024, 600, webview.HintNone)
	w.Navigate(fmt.Sprintf("http://localhost:%d", port))
	w.Run()
}
