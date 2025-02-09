package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Anton-Kraev/vk-test-assignment/backend/internal/http/handler"
)

func Start(handler handler.Handler) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/containers", handler.GetContainers)
	e.PATCH("/containers", handler.UpdateContainersPing)

	return e.Start(":8080")
}
