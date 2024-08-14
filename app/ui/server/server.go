package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	port string
}

func NewServer(port string) *Server {
	return &Server{
		echo: echo.New(),
		port: port,
	}
}

func (s *Server) Start() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	go s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf(":%s", s.port)))
}

func (s *Server) Stop() {
	go s.echo.Shutdown(context.Background())
}
