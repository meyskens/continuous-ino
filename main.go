package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Contininuous Ino -- CI for Arduino")
	fmt.Println("==================================")

	e := echo.New()
	e.GET("/", serveRoot)
	e.Logger.Fatal(e.Start(":80"))
}
