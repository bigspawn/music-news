package internal

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-pkgz/lgr"
)

const (
	driver      = "sqlite3"
	insertQuery = `			INSERT INTO main.news(title, playlist, date_time, imageurl, downloadurl, pageurl, posted, created_at) 
							VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	selectNotified = `		SELECT id, title, playlist, imageurl, date_time, downloadurl
							FROM main.news
							WHERE not notified
							  AND title NOT LIKE '%Single%'
							  AND title NOT LIKE '%single%'
							  AND created_at > strftime('%Y-%m-%d %H:%M:%S+00:00', 'now', 'utc', '-1 months')
							ORDER BY date_time`

	updateNotified = `		UPDATE main.news 
							SET notified = true 
							WHERE id = $1`

	selectUnpublished = `	SELECT id, title, playlist, imageurl, date_time, downloadurl, pageurl
							FROM main.news
							WHERE posted = false
							  AND created_at > now() - interval '4 week'
							  AND created_at < now() - interval '1 minute'
							ORDER BY created_at`

	updatePosted = `		UPDATE main.news 
							SET posted = true 
							WHERE title = $1`

	selectExists = `		SELECT * 
							FROM main.news 
							WHERE title LIKE '%' || $1 || '%'`
)

// News is an article structure.
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
	db  *sql.DB
	lgr lgr.L
}

func (s *Store) Exist(ctx context.Context, title string) (bool, error) {
	row, err := s.db.QueryContext(ctx, selectExists, title)
	if err != nil {
		return false, err
	}
	defer func() { _ = row.Close() }()

	if row != nil && row.Next() {
		return true, nil
	}
	return false, nil
}

func (s *Store) Insert(ctx context.Context, n News) error {
	var userID int

	row := s.db.QueryRowContext(ctx, insertQuery,
		n.Title, n.Text, n.DateTime, n.ImageLink, n.DownloadLink[0], n.PageLink, false, time.Now())

	if err := row.Scan(&userID); err != nil {
		return err
	}

	s.lgr.Logf("[DEBUG] insert %v", n)

	return nil
}

func (s *Store) GetWithNotifyFlag(ctx context.Context) ([]News, error) {
	rows, err := s.db.QueryContext(ctx, selectNotified)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var (
		ID           = new(int)
		Title        = new(string)
		Text         = new(string)
		ImageLink    = new(string)
		DownloadLink = new(string)
		DateTime     = new(time.Time)
		notifies     = make([]News, 0)
	)
	for rows.Next() {
		if err := rows.Scan(ID, Title, Text, ImageLink, DateTime, DownloadLink); err != nil {
			return nil, err
		}
		notifies = append(notifies, News{
			ID:           *ID,
			Title:        *Title,
			Text:         *Text,
			ImageLink:    *ImageLink,
			DownloadLink: strings.Split(*DownloadLink, ","),
			DateTime:     DateTime,
		})
	}
	return notifies, nil
}

func (s *Store) UpdateNotifyFlag(ctx context.Context, item News) error {
	result, err := s.db.ExecContext(ctx, updateNotified, item.ID)
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

func (s *Store) GetUnpublished(ctx context.Context) (map[string]News, error) {
	rows, err := s.db.QueryContext(ctx, selectUnpublished)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var (
		ID           = new(int)
		Title        = new(string)
		Text         = new(string)
		ImageLink    = new(string)
		PageLink     = new(string)
		DownloadLink = new(string)
		DateTime     = new(time.Time)
		unpublished  = make(map[string]News)
	)
	for rows.Next() {
		if err := rows.Scan(ID, Title, Text, ImageLink, DateTime, DownloadLink, PageLink); err != nil {
			return nil, err
		}
		n := News{
			ID:           *ID,
			Title:        *Title,
			Text:         *Text,
			ImageLink:    *ImageLink,
			DownloadLink: strings.Split(*DownloadLink, ","),
			DateTime:     DateTime,
			PageLink:     *PageLink,
		}
		unpublished[n.Title] = n
	}
	return unpublished, nil
}

func (s *Store) SetPosted(ctx context.Context, title string) error {
	result, err := s.db.ExecContext(ctx, updatePosted, title)
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
