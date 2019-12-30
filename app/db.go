package main

import (
	"database/sql"
	"log"
	"time"
)

var (
	insertQuery = `INSERT INTO public.news(title, playlist, date_time, imageurl, downloadurl) 
							   VALUES ($1, $2, $3, $4, $5) RETURNING id`
	driver = "postgres"
)

// News is an article structure
type News struct {
	ID           int
	Title        string
	Text         string
	ImageLink    string
	DownloadLink []string
	PageLink     string
	DateTime     *time.Time
}

type NewsStore struct {
	conn *sql.DB
}

func NewNewsStore(conn string) (*NewsStore, error) {
	db, err := sql.Open(driver, conn)
	if err != nil {
		return nil, err
	}
	return &NewsStore{conn: db}, nil
}

// fixme: change on one operation UPDATE
func (nw *NewsStore) Exist(title string) (bool, error) {
	row, err := nw.conn.Query("SELECT * FROM public.news WHERE title LIKE '%' || $1 || '%'", title)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = row.Close()
	}()
	if row != nil && row.Next() {
		return true, nil
	}
	return false, nil
}

func (nw *NewsStore) Insert(n *News) error {
	var userid int
	row := nw.conn.QueryRow(insertQuery, n.Title, n.Text, n.DateTime, n.ImageLink, n.DownloadLink[0])
	err := row.Scan(&userid)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] insert %v", n)
	return nil
}
