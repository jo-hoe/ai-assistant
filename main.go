package main

import (
	"fmt"

	"github.com/jo-hoe/ai-assistent/app/server"
	webview "github.com/webview/webview_go"
)

func main() {
	port := "8080"
	server := server.NewServer("8080")
	go server.Start()
	defer server.Stop()

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("AI Assistant")
	w.SetSize(1024, 600, webview.HintNone)
	w.Navigate(fmt.Sprintf("http://localhost:%s", port))
	w.Run()
}
