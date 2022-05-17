package repository

import (
	"database/sql"
	"log"
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
	query := "select s.timeStamp  from stories s order by s.timeStamp DESC  LIMIT 1"
	var tmstmp int64
	rp.db.QueryRow(query).Scan(&tmstmp)
	return time.UnixMilli(tmstmp)
}

func (rp *Repository) GetStories() []story.Story {
	query := "select s.storyId,s.title,s.score  from stories s"
	rows, err := rp.db.Query(query)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	stories := []story.Story{}

	for rows.Next() {
		st := story.Story{}
		if err := rows.Scan(&st.Id, &st.Title, &st.Score); err != nil {
			log.Println(err)
		}
		stories = append(stories, st)
	}

	return stories
}

//GetLastTimeStamp()

//GetStories()
