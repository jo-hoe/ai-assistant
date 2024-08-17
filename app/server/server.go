package server

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())

	return &Server{
		echo: e,
	}
}

type Counter struct {
	Count int
}

var count Counter

func (s *Server) Start(port int) {
	s.echo.Renderer = &Template{
		templates: template.Must(template.ParseFS(templateFS, viewsPattern)),
	}

	count = Counter{Count: 0}

	s.echo.GET("/", IndexHandler)
	s.echo.POST("/count", CountUp)

	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf("127.0.0.1:%d", port)))
}

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", count)
}

func CountUp(c echo.Context) error {
	count.Count++
	return c.Render(http.StatusOK, "count", count)
}

func (s *Server) Stop() {
	err := s.echo.Shutdown(context.Background())
	if err != nil {
		s.echo.Logger.Fatal(err)
	}
}
