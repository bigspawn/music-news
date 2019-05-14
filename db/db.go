package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type News struct {
	ID           int
	Title        string
	Text         string
	ImageLink    string
	DownloadLink []string
	PageLink     string
	DateTime     *time.Time
}

func Connection(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Select(db *sql.DB, title string) (*sql.Rows, error) {
	return db.Query("SELECT * FROM public.news WHERE title LIKE '%' || $1 || '%'", title)
}

func Insert(db *sql.DB, n News) {
	db.QueryRow("INSERT INTO public.news(title, playlist, date_time, imageurl, downloadurl) VALUES ($1, $2, $3, $4, $5)",
		n.Title, n.Text, n.DateTime, n.ImageLink, n.DownloadLink[0])
}
