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

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server := server.NewServer()
	defer server.Stop()

	if *headless {
		server.Start(config.Port)
	} else {
		go server.Start(config.Port)
		cmd.StartWebview(config.Port)
	}
}
