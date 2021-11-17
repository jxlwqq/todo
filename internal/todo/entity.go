package todo

import "time"

type Item struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	RemindAt    time.Time `json:"remind_at"`
}
