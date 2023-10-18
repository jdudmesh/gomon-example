package console

import (
	"time"
)

type LogRun struct {
	ID        int64     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type LogEvent struct {
	ID        int64     `db:"id" json:"id"`
	RunID     int64     `db:"run_id" json:"runId"`
	EventType string    `db:"event_type" json:"eventType"`
	EventData string    `db:"event_data" json:"eventData"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
