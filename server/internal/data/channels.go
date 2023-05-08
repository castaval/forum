package data

import (
	"forum/internal/validator"
	"time"
)

type Channel struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Version   int32     `json:"version"`
}

func ValidateChannel(v *validator.Validator, channel *Channel) {
	v.Check(channel.Title != "", "title", "must be provided")
	v.Check(len(channel.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(channel.UserID != 0, "user_id", "must be provided")
	v.Check(channel.UserID > 0, "user_id", "must be a positive integer")
}
