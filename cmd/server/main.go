package main

import (
	"net/http"

	client "github.com/jdudmesh/gomon-client"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	log.Info("starting server")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t, err := client.NewEcho("./views/*html")
	if err != nil {
		log.Fatal(err)
	}
	defer t.Close()

	go func() {
		err := t.ListenAndServe()
		if err != nil {
			log.Error(err)
		}
	}()

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "World")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
