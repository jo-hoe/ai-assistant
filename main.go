package main

import (
	"flag"

	"github.com/jo-hoe/ai-assistent/app/cmd"
	"github.com/jo-hoe/ai-assistent/app/config"
	"github.com/jo-hoe/ai-assistent/app/server"
)

func main() {
	headless := flag.Bool("headless", false, "If true the webview will not be started.")
	flag.Parse()

	config := config.GetConfig()
	server := server.NewServer(config)
	defer server.Stop()

	if *headless {
		server.Start(config.Port)
	} else {
		go server.Start(config.Port)
		cmd.StartWebview(config.Port)
	}
}
