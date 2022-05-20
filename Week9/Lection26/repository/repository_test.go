package repository

import (
	"database/sql"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

const (
	createStoryTable = "CREATE table if not exists stories(storyId int primary key, title text, score int, timeStamp datetime default current_timestamp)"
	insertStory      = "insert into stories (storyId,title,score) values(?,?,?,?)"
)

func TestLastStoryTimeStamp(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Fail to open in-memory DB")
	}
	//migrate/populate the schema

	_, err = mockDb.Exec(createStoryTable)
	if err != nil {
		t.Fatal("Fail to create table")
	}

	repo := NewRepository(mockDb)
	result := repo.GetLastStoryTimeStamp()

	if result == time.Now() {
		t.Fatal("Fail to create initial condition")
	}

	wantedTime := time.Now().Add(time.Hour)
	mockDb.Exec(insertStory, 0, "Unit test", 15, wantedTime)

	mockDb.Exec(insertStory, 1, "Unit test1", 15, time.Now().Add(-time.Hour))

	result = repo.GetLastStoryTimeStamp()
	if result == wantedTime {
		t.Fatal("Fail to take latest timeStamp")
	}
}
