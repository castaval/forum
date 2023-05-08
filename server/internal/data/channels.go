package data

import "time"

type Channel struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Version   int32     `json:"version"`
}
