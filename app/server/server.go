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

func (s *Server) Start() {
	s.echo.Renderer = &Template{
		templates: template.Must(template.ParseFS(templateFS, viewsPattern)),
	}

	s.echo.GET("/", HomeHandler)

	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf(":%s", s.port)))
}

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (s *Server) Stop() {
	s.echo.Shutdown(context.Background())
}
