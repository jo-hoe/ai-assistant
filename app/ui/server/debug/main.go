package main

import "github.com/jo-hoe/ai-assistent/app/ui/server"

func main() {
	server := server.NewServer("8080")
	server.Start()
	defer server.Stop()
}
