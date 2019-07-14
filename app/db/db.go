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
	log.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Open connection")
	return db
}

func Select(db *sql.DB, title string) (*sql.Rows, error) {
	return db.Query("SELECT * FROM public.news WHERE title LIKE '%' || $1 || '%'", title)
}

// fixme: переделать на запрос с проверкой все ли успешно
func Insert(db *sql.DB, n News) error {
	log.Printf("[DEBUG] insert %v", n)
	var userid int
	return db.QueryRow(`INSERT INTO public.news(title, playlist, date_time, imageurl, downloadurl) 
							   VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		n.Title, n.Text, n.DateTime, n.ImageLink, n.DownloadLink[0]).Scan(&userid)
}