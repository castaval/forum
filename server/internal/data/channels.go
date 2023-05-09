package data

import (
	"database/sql"
	"errors"
	"forum/internal/validator"
	"time"
)

type Channel struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Version   int32     `json:"version"`
}

func ValidateChannel(v *validator.Validator, channel *Channel) {
	v.Check(channel.Title != "", "title", "must be provided")
	v.Check(len(channel.Title) <= 500, "title", "must not be more than 500 bytes long")
}

type ChannelModel struct {
	DB *sql.DB
}

func (c ChannelModel) Insert(channel *Channel) error {
	query := `
		INSERT INTO channels (title)
		VALUES ($1)
		RETURNING id, created_at, version
	`

	args := []interface{}{channel.Title}

	return c.DB.QueryRow(query, args...).Scan(&channel.ID, &channel.CreatedAt, &channel.Version)
}

func (c ChannelModel) Get(id int64) (*Channel, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, title, version
		FROM channels
		WHERE id = $1`

	var channel Channel
	err := c.DB.QueryRow(query, id).Scan(
		&channel.ID,
		&channel.CreatedAt,
		&channel.Title,
		&channel.Version,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &channel, nil
}

func (c ChannelModel) Update(channel *Channel) error {
	query := `
		UPDATE channels
		SET title = $1, version = version + 1
		WHERE id = $2
		RETURNING version
	`

	args := []interface{}{
		channel.Title,
		channel.ID,
	}

	return c.DB.QueryRow(query, args...).Scan(&channel.Version)
}

func (c ChannelModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM channels
		WHERE id = $1
	`

	result, err := c.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
