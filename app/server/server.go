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

type Counter struct {
	Count int
}

var count Counter

func (s *Server) Start() {
	s.echo.Renderer = &Template{
		templates: template.Must(template.ParseFS(templateFS, viewsPattern)),
	}

	count = Counter{Count: 0}

	s.echo.GET("/", IndexHandler)
	s.echo.POST("/count", CountUp)

	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf(":%s", s.port)))
}

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", count)
}

func CountUp(c echo.Context) error {
	count.Count++
	return c.Render(http.StatusOK, "index.html", count)
}

func (s *Server) Stop() {
	s.echo.Shutdown(context.Background())
}
