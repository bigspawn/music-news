package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	driver      = "postgres"
	insertQuery = `		INSERT INTO public.news(title, playlist, date_time, imageurl, downloadurl) 
						VALUES ($1, $2, $3, $4, $5) RETURNING id`

	selectNotified = `	SELECT id, title, playlist, imageurl, date_time, downloadurl
						FROM public.news
						WHERE notified = false
						  AND title NOT LIKE '%Single%'
						  AND title NOT LIKE '%single%'
						  AND date_time > now() - interval '1 week'
						ORDER BY date_time`

	updateNotified = `	UPDATE news 
						SET notified = true 
						WHERE id = $1`
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

type Store struct {
	conn *sql.DB
}

func NewNewsStore(conn string) (*Store, error) {
	db, err := sql.Open(driver, conn)
	if err != nil {
		return nil, err
	}
	return &Store{conn: db}, nil
}

// fixme: change on one operation UPDATE
func (s *Store) Exist(title string) (bool, error) {
	row, err := s.conn.Query("SELECT * FROM public.news WHERE title LIKE '%' || $1 || '%'", title)
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

func (s *Store) Insert(n *News) error {
	var userid int
	row := s.conn.QueryRow(insertQuery, n.Title, n.Text, n.DateTime, n.ImageLink, n.DownloadLink[0])
	err := row.Scan(&userid)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] insert %v", n)
	return nil
}

func (s *Store) GetWithNotifyFlag(ctx context.Context) ([]*News, error) {
	forNotify := make([]*News, 0)

	rows, err := s.conn.QueryContext(ctx, selectNotified)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			ID           = new(int)
			Title        = new(string)
			Text         = new(string)
			ImageLink    = new(string)
			DownloadLink = new(string)
			DateTime     = new(time.Time)
		)
		if err := rows.Scan(ID, Title, Text, ImageLink, DateTime, DownloadLink); err != nil {
			return nil, err
		}
		forNotify = append(forNotify, &News{
			ID:           *ID,
			Title:        *Title,
			Text:         *Text,
			ImageLink:    *ImageLink,
			DownloadLink: strings.Split(*DownloadLink, ","),
			DateTime:     DateTime,
		})
	}
	return forNotify, nil
}

func (s *Store) UpdateNotifyFlag(ctx context.Context, item *News) error {
	result, err := s.conn.ExecContext(ctx, updateNotified, item.ID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}
