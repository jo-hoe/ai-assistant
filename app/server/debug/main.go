package main

import "github.com/jo-hoe/ai-assistent/app/server"

func main() {
	server := server.NewServer("8080")
	defer server.Stop()
	server.Start()
}
