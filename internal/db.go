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
							  AND created_at > strftime('%Y-%m-%d %H:%M:%S+00:00', 'now', 'utc', '-1 months')
							ORDER BY created_at`

	updatePosted = `		UPDATE main.news 
							SET posted = true 
							WHERE title = $1`

	selectExists = `		SELECT * 
							FROM main.news 
							WHERE title LIKE '%' || $1 || '%'`

	selectAll = `			SELECT id, title, date_time, downloadurl, imageurl, playlist, posted, notified, created_at
							FROM main.news`

	updatePostedByID = `	UPDATE main.news 
							SET posted = true 
							WHERE id = $1`

	updatePostedAndNotifiedByID = `	UPDATE main.news 
									SET posted = true, notified = true 
									WHERE id = $1`
)

// News is an article structure.
type News struct {
	ID           int
	Title        string
	DateTime     time.Time
	DownloadLink []string
	ImageLink    string
	PageLink     string
	Text         string
	Posted       bool
	Notified     bool
	CreatedAt    time.Time
	TitleHash    string
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
	var id int
	row := s.db.QueryRowContext(ctx, insertQuery, n.Title, n.Text, n.DateTime, n.ImageLink, n.DownloadLink[0],
		n.PageLink, false, time.Now().UTC())
	return row.Scan(&id)
}

func (s *Store) GetWithNotifyFlag(ctx context.Context) ([]News, error) {
	rows, err := s.db.QueryContext(ctx, selectNotified)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var (
		id           = new(int)
		title        = new(string)
		text         = new(string)
		imageLink    = new(string)
		downloadLink = new(string)
		dateTime     = new(time.Time)
		notifies     = make([]News, 0)
	)
	for rows.Next() {
		if err := rows.Scan(id, title, text, imageLink, dateTime, downloadLink); err != nil {
			return nil, err
		}
		notifies = append(notifies, News{
			ID:           *id,
			Title:        *title,
			Text:         *text,
			ImageLink:    *imageLink,
			DownloadLink: strings.Split(*downloadLink, ","),
			DateTime:     *dateTime,
		})
	}
	return notifies, nil
}

func (s *Store) UpdateNotifyFlag(ctx context.Context, item News) error {
	return s.exec(ctx, updateNotified, item.ID)
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
			DateTime:     *DateTime,
			PageLink:     *PageLink,
		}
		unpublished[n.Title] = n
	}
	return unpublished, nil
}

func (s *Store) SetPosted(ctx context.Context, title string) error {
	return s.exec(ctx, updatePosted, title)
}

func (s *Store) SetPostedByID(ctx context.Context, id int) error {
	return s.exec(ctx, updatePostedAndNotifiedByID, id)
}

func (s *Store) GetAll(ctx context.Context) ([]News, error) {
	rows, err := s.db.QueryContext(ctx, selectAll)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	items := make([]News, 0)
	for rows.Next() {
		var n News
		var links string
		err := rows.Scan(
			&n.ID,
			&n.Title,
			&n.DateTime,
			&links,
			&n.ImageLink,
			&n.Text,
			&n.Posted,
			&n.Notified,
			&n.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		n.DownloadLink = strings.Split(links, ",")
		items = append(items, n)
	}

	return items, nil
}

func (s *Store) SetPostedAndNotified(ctx context.Context, id int) error {
	return s.exec(ctx, updatePostedByID, id)
}

func (s *Store) exec(ctx context.Context, sql string, args ...interface{}) error {
	result, err := s.db.ExecContext(ctx, sql, args)
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
