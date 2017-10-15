package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/labstack/echo"
)

func serveRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Continuous Ino Server")
}

func serveWebhook(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.String(http.StatusBadRequest, "I only support the JSON format")
	}
	payload, err := github.ValidatePayload(c.Request(), []byte(cfg.GitHub.WebhookSecret))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid secret")
	}

	info := github.WebHookPayload{}
	json.Unmarshal(payload, &info)

	handleHookPayload(info)

	return c.String(http.StatusOK, "OK")
}
