package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/validator"
	"time"
)

type Thread struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      int64     `json:"user_id"`
	ChannelID   int64     `json:"channel_id"`
	CreatedAt   time.Time `json:"-"`
	Version     int32     `json:"version"`
}

func ValidateThreads(v *validator.Validator, thread *Thread) {
	v.Check(thread.Title != "", "title", "must be provided")
	v.Check(len(thread.Title) <= 500, "title", "must not be more than 500 long")
	v.Check(thread.Description != "", "description", "must be provided")
	v.Check(len(thread.Description) >= 10 && len(thread.Description) <= 5000, "description", "must be in range 10-5000 long")
	v.Check(thread.UserID != 0, "user_id", "must be provided")
	v.Check(thread.UserID > 0, "user_id", "must be positive number")
	v.Check(thread.ChannelID != 0, "channel_id", "must be provided")
	v.Check(thread.ChannelID > 0, "channel_id", "must be positive number")
}

type ThreadModel struct {
	DB *sql.DB
}

func (t ThreadModel) Insert(thread *Thread) error {
	query := `
		INSERT INTO threads (title, description, user_id, channel_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version
	`

	args := []interface{}{thread.Title, thread.Description, thread.UserID, thread.ChannelID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return t.DB.QueryRowContext(ctx, query, args...).Scan(&thread.ID, &thread.CreatedAt, &thread.Version)
}

func (t ThreadModel) Get(id int64) (*Thread, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, title, description, user_id, channel_id, created_at, version
		FROM threads
		WHERE id = $1
	`

	var thread Thread

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, id).Scan(
		&thread.ID,
		&thread.Title,
		&thread.Description,
		&thread.UserID,
		&thread.ChannelID,
		&thread.CreatedAt,
		&thread.Version,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &thread, nil
}

func (t ThreadModel) Update(thread *Thread) error {
	query := `
		UPDATE threads
		SET title = $1, description = $2, version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING version
	`

	args := []interface{}{
		thread.Title,
		thread.Description,
		thread.ID,
		thread.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, args...).Scan(&thread.Version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrEditConflict
		}
		return err
	}

	return nil
}

func (t ThreadModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE from threads
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := t.DB.ExecContext(ctx, query, id)
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

func (t ThreadModel) GetAll(filters Filters) ([]*Thread, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, title, description, user_id, channel_id, created_at, version
		FROM threads
		ORDER by %s %s, id ASC
		LIMIT $1 OFFSET $2`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{filters.limit(), filters.offset()}

	rows, err := t.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	threads := []*Thread{}

	for rows.Next() {
		var thread Thread

		err := rows.Scan(
			&totalRecords,
			&thread.ID,
			&thread.Title,
			&thread.Description,
			&thread.UserID,
			&thread.ChannelID,
			&thread.CreatedAt,
			&thread.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		threads = append(threads, &thread)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return threads, metadata, nil
}
