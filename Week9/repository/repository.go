package repository

import (
	"database/sql"
	"main/story"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (rp *Repository) GetLastStoryTimeStamp() time.Time {
	return time.Now()
}

func (rp *Repository) GetStories() []story.Story {
	return []story.Story{}
}

//GetLastTimeStamp()

//GetStories()
