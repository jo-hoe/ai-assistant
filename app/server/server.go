package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/jo-hoe/ai-assistent/app/aiclient"
	"github.com/jo-hoe/ai-assistent/app/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	e := echo.New()
	e.Use(middleware.Logger())

	return &Server{
		echo:   e,
		config: config,
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

func (s *Server) Stop() {
	err := s.echo.Shutdown(context.Background())
	if err != nil {
		s.echo.Logger.Fatal(err)
	}
}

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

type question struct {
	Question string `form:"question"`
}

type Conversation struct {
	Conversation []QnA
}

type QnA struct {
	Question string
	Answer   string
}

func AskAIHandler(c echo.Context) error {
	prefix := c.QueryParam("question-prefix")

	var question question
	if err := c.Bind(&question); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	config := config.GetConfig()

	responseChannel, err := config.AIClients.GetAnswer([]aiclient.Message{
		{
			Role:    "user",
			Content: fmt.Sprintf("%s %s", prefix, question.Question),
		},
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	stringBuilder := strings.Builder{}
	for answerPart := range responseChannel {
		stringBuilder.WriteString(answerPart.Answer)
	}
	answer := stringBuilder.String()

	return c.Render(http.StatusOK, "qna", QnA{
		Answer:   answer,
		Question: question.Question,
	})
}
