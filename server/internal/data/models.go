package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Users    UserModel
	Channels ChannelModel
	Threads  ThreadModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Channels: ChannelModel{DB: db},
		Threads:  ThreadModel{DB: db},
		Users:    UserModel{DB: db},
	}
}
