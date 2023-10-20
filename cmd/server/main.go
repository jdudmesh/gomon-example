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

	http.HandleFunc("/run/events", func(w http.ResponseWriter, r *http.Request) {
		var err error
		run := r.URL.Query().Get("run")
		stm := r.URL.Query().Get("stm")
		filter := r.URL.Query().Get("filter")
		events := []console.LogEvent{}

		params := map[string]interface{}{"run_id": run}
		sql := "SELECT * FROM events WHERE run_id = :run_id "
		if !(stm == "" || stm == "all") {
			sql += " AND event_type = :event_type "
			params["event_type"] = stm
		}
		if filter != "" {
			sql += " AND event_data LIKE :event_data "
			params["event_data"] = "%" + filter + "%"
		}
		sql += " ORDER BY created_at ASC;"

		res, err := db.NamedQuery(sql, params)
		if err != nil {
			log.Errorf("getting event: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Close()
		for res.Next() {
			var ev console.LogEvent
			err = res.StructScan(&ev)
			if err != nil {
				log.Errorf("scanning event: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			events = append(events, ev)
		}

		if len(events) == 0 {
			_, err = w.Write([]byte("<div class=\"text-2xl text-bold\">no events found</div>"))
			if err != nil {
				log.Errorf("writing response: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		markup := console.EventList(events)
		err = markup.Render(r.Context(), w)
		if err != nil {
			log.Errorf("rendering index: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
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
