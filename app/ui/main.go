package main

import webview "github.com/webview/webview_go"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("AI Assistant")
	w.SetSize(480, 320, webview.HintNone)
	w.SetHtml("Hello I'm your Assistant")
	w.Run()
}