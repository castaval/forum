package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.DB.QueryRowContext(ctx, query, args...).Scan(&channel.ID, &channel.CreatedAt, &channel.Version)
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, id).Scan(
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.DB.QueryRowContext(ctx, query, args...).Scan(&channel.Version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrEditConflict
		}
		return err
	}

	return nil
}

func (c ChannelModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM channels
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := c.DB.ExecContext(ctx, query, id)
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

func (c ChannelModel) GetAll(filters Filters) ([]*Channel, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, created_at, title, version
		FROM channels
		ORDER by %s %s, id ASC
		LIMIT $1 OFFSET $2`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{filters.limit(), filters.offset()}

	rows, err := c.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	channels := []*Channel{}

	for rows.Next() {
		var channel Channel

		err := rows.Scan(
			&totalRecords,
			&channel.ID,
			&channel.CreatedAt,
			&channel.Title,
			&channel.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		channels = append(channels, &channel)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return channels, metadata, nil
}
