package main

import (
	"fmt"
	"net/http"
	"os"

	client "github.com/jdudmesh/gomon-client"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS runs (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_runs_created_at ON runs(created_at);

CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	run_id INTEGER NOT NULL,
	event_type TEXT NOT NULL,
	event_data TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_events_event_type ON events(event_type);
CREATE INDEX IF NOT EXISTS idx_events_created_at ON events(created_at);
`

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/assets", "./static")

	db, err := sqlx.Connect("sqlite3", "./.gomon/gomon.db")
	if err != nil {
		log.Fatalf("connecting to sqlite: %v", err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("creating console capture db schema: %v", err)
	}

	t, err := client.NewEcho("views/*.html", e.Logger)
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
		c.Logger().Info("Hello, World!")
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/other", func(c echo.Context) error {
		c.Logger().Info("Other!")
		return c.Render(http.StatusOK, "other.html", nil)
	})

	if p, ok := os.LookupEnv("PORT"); ok {
		e.Logger.Fatal(e.Start(":" + p))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}
}
