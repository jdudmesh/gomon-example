package main

import (
	"fmt"
	"net/http"
	"os"

	templates "github.com/jdudmesh/gomon-client"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/assets", "./static")

	t, err := templates.NewEcho("views/*.html", e.Logger)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer t.Close()
	if err := t.Run(); err != nil {
		panic(err)
	}

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/other", func(c echo.Context) error {
		return c.Render(http.StatusOK, "other.html", nil)
	})

	if p, ok := os.LookupEnv("PORT"); ok {
		e.Logger.Fatal(e.Start(":" + p))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}
}
