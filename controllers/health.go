package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message"`
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, Message{"Ok"})
}
