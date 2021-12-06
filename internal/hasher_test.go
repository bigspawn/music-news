package internal

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/go-pkgz/lgr"
	_ "github.com/mattn/go-sqlite3"
)

func TestName(t *testing.T) {
	path := "file:/Users/bigspawn/Developer/projects/github.com/bigspawn/music-news/music-news-db/music_news_prod.sqlite"

	db, err := sql.Open(driver, path)
	if err != nil {
		t.Fatal(err)
	}

	s := Store{
		db:  db,
		lgr: lgr.Default(),
	}

	items, err := s.GetAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	reg, err := regexp.Compile(`\w+`)
	if err != nil {
		t.Fatal(err)
	}

	for i := range items {
		find := reg.FindAllString(items[i].Title, -1)

		h := sha256.New()
		for i := range find {
			_, err := io.WriteString(h, strings.ToLower(find[i]))
			if err != nil {
				t.Fatal(err)
			}
		}

		items[i].TitleHash = hex.EncodeToString(h.Sum(nil))
	}

	err = s.SetTitleHash(context.Background(), items)
	if err != nil {
		t.Fatal(err)
	}

}
