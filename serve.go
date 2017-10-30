package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/labstack/echo"
)

func serveWebhook(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.String(http.StatusBadRequest, "I only support the JSON format")
	}
	payload, err := github.ValidatePayload(c.Request(), []byte(cfg.GitHub.WebhookSecret))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusUnauthorized, "Invalid secret")
	}

	info := github.WebHookPayload{}
	json.Unmarshal(payload, &info)

	handleHookPayload(github.WebHookType(c.Request()), info)

	return c.String(http.StatusOK, "OK")
}
