package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func serveGetBuild(c echo.Context) error {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "invalid ID"})
	}
	run, err := store.GetRun(ID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "unable to get build"})
	}

	return c.JSON(http.StatusOK, run)
}

func serveGetAllBuilds(c echo.Context) error {
	return c.JSON(http.StatusOK, store.GetAllRuns())
}
