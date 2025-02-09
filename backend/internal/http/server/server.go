package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Anton-Kraev/vk-test-assignment/backend/internal/http/handler"
)

const (
	frontendURL = "http://localhost:3000"
	pingerURL   = "http://localhost:5000"
	port        = ":8080"
)

func Start(handler handler.Handler) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontendURL, pingerURL},
		AllowMethods: []string{echo.GET, echo.PATCH},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/containers", handler.GetContainers)
	e.PATCH("/containers", handler.UpdateContainersPing)

	return e.Start(port)
}
