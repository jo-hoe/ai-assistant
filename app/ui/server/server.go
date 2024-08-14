package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jo-hoe/ai-assistent/app/ui/server/views"

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

	s.echo.GET("/", HomeHandler)
	s.echo.Logger.Fatal(s.echo.Start(fmt.Sprintf(":%s", s.port)))
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, views.Home())
}

func (s *Server) Stop() {
	s.echo.Shutdown(context.Background())
}
