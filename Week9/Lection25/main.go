package main

import (
	"database/sql"
	"log"
	handlers "main/handlers"
	repository "main/repository"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	router := http.NewServeMux()
	db, err := sql.Open("sqlite", "dbHN.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	router.Handle("/api/top", handlers.HandleTopStories(repo)) //return top 10
	//router.Handle("top", nil)                    //view html template
	//log.Println("started")
	http.ListenAndServe(":8000", router)
}
