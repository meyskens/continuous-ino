package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func serveRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Continuous Ino Server")
}
