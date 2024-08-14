package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	port string
}

func NewServer(port string) *Server {
	e := echo.New()
	e.Use(middleware.Logger())

	return &Server{
		echo: e,
		port: port,
	}
}

type Count struct {
	Count int `json:"count"`
}

func (s *Server) Start() {
	s.echo.Renderer = NewTemplates()

	count := &Count{Count: 0}
	s.echo.GET("/", func(c echo.Context) error {
		count.Count++
		return c.Render(http.StatusOK, "index", count)
	})
	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf(":%s", s.port)))
}

func (s *Server) Stop() {
	s.echo.Shutdown(context.Background())
}
