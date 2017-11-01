package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/meyskens/continuous-ino/config"
	"github.com/meyskens/continuous-ino/storage"
)

var cfg config.Configuration
var db *bolt.DB
var store storage.Storage

func main() {
	fmt.Println("Contininuous Ino -- CI for Arduino")
	fmt.Println("==================================")

	cfg = config.GetConfiguration()
	setUpGitHub()

	db, err := bolt.Open(cfg.Database.Path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store = storage.New(db)

	e := echo.New()
	e.Use(middleware.CORS())

	e.POST("/webhook", serveWebhook)
	e.GET("/api/build/:id", serveGetBuild)
	e.GET("/api/build/all", serveGetAllBuilds)

	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":80"))
}
