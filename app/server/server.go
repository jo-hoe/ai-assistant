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

func (s *Server) Start(port int) {
	s.echo.Renderer = &Template{
		templates: template.Must(template.ParseFS(templateFS, viewsPattern)),
	}

	s.echo.GET("/", IndexHandler)
	s.echo.POST("/ask", AskAIHandler)

	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf("127.0.0.1:%d", port)))
}

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func AskAIHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "answer", nil)
}

func (s *Server) Stop() {
	err := s.echo.Shutdown(context.Background())
	if err != nil {
		s.echo.Logger.Fatal(err)
	}
}
