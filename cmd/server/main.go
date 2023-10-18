package main

import (
	"gohtmx/console"
	"net/http"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting server")

	db, err := sqlx.Connect("sqlite3", "./.gomon/gomon.db")
	if err != nil {
		log.Fatalf("connecting to sqlite: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("request: ", r.URL.Path)

		runs := []console.LogRun{}
		err = db.Select(&runs, "SELECT * FROM runs ORDER BY created_at DESC")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		currentRun := int(runs[0].ID)

		run := int(runs[0].ID)
		runParam := r.URL.Query().Get("run")
		if !(runParam == "" || runParam == "current") {
			run, err = strconv.Atoi(runParam)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		events := []console.LogEvent{}
		err = db.Select(&events, "SELECT * FROM events WHERE run_id = ? ORDER BY created_at ASC", run)
		if err != nil {
			log.Errorf("getting event: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		page := console.Console(currentRun, runs, events)
		err = page.Render(r.Context(), w)
		if err != nil {
			log.Errorf("rendering index: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		log.Info("request: ", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	})

	host := ":8080"
	if p, ok := os.LookupEnv("PORT"); ok {
		host = ":" + p
	}
	log.Infof("listening on %s", host)
	log.Fatal(http.ListenAndServe(host, nil))
}
