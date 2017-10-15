package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/meyskens/continuous-ino/config"
)

var cfg config.Configuration

func main() {
	fmt.Println("Contininuous Ino -- CI for Arduino")
	fmt.Println("==================================")

	cfg = config.GetConfiguration()
	setUpGitHub()

	e := echo.New()
	e.GET("/", serveRoot)
	e.POST("/webhook", serveWebhook)
	e.Logger.Fatal(e.Start(":80"))
}
