package main

import (
	"flag"
	"log"

	"github.com/jo-hoe/ai-assistent/app/cmd"
	"github.com/jo-hoe/ai-assistent/app/server"
)

func main() {
	port := flag.Int("port", -1, "Select to port number where the server will be running.")
	headless := flag.Bool("headless", false, "If true the webview will not be started.")
	flag.Parse()

	if *port == -1 {
		freePort, err := server.GetFreeTcpPort()
		if err != nil {
			log.Fatal(err)
		}
		*port = freePort
	}
	log.Printf("starting server on port :%d", *port)

	server := server.NewServer()
	defer server.Stop()

	if *headless {
		server.Start(*port)
	} else {
		go server.Start(*port)
		cmd.StartWebview(*port)
	}
}
