package internal

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-pkgz/lgr"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func newTestStore(t *testing.T) *Store {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)

	_, err = db.Exec(`CREATE TABLE news (
		id          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		title       VARCHAR(255),
		date_time   TIMESTAMP,
		downloadurl TEXT,
		imageurl    TEXT,
		pageurl     VARCHAR(255),
		playlist    TEXT,
		posted      BOOLEAN NOT NULL DEFAULT true,
		notified    BOOL NOT NULL DEFAULT false,
		created_at  TIMESTAMP
	)`)
	require.NoError(t, err)

	store, err := NewStore(StoreParams{Lgr: lgr.NoOp, DB: db})
	require.NoError(t, err)

	t.Cleanup(func() { _ = db.Close() })
	return store
}

func insertTestNews(t *testing.T, store *Store, title string) int {
	t.Helper()
	id, err := store.Insert(context.Background(), News{
		Title:    title,
		DateTime: time.Now(),
	})
	require.NoError(t, err)
	return id
}

func TestStore_IncrementNotifyAttempts(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()
	id := insertTestNews(t, store, "Test Band - Test Album (2026)")

	// first increment: 0 → 1
	err := store.IncrementNotifyAttempts(ctx, id)
	require.NoError(t, err)

	// verify state
	var attempts int
	var nextRetry sql.NullTime
	err = store.DB.QueryRowContext(ctx,
		"SELECT notify_attempts, notify_next_retry FROM news WHERE id = $1", id,
	).Scan(&attempts, &nextRetry)
	require.NoError(t, err)
	require.Equal(t, 1, attempts)
	require.True(t, nextRetry.Valid, "notify_next_retry should be set")
}

func TestStore_IncrementNotifyAttempts_MarksNotifiedAtMax(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()
	id := insertTestNews(t, store, "Test Band - Max Attempts (2026)")

	// set attempts to maxNotifyAttempts - 1
	_, err := store.DB.ExecContext(ctx,
		"UPDATE news SET notify_attempts = $1 WHERE id = $2", maxNotifyAttempts-1, id)
	require.NoError(t, err)

	// this increment should mark notified = true
	err = store.IncrementNotifyAttempts(ctx, id)
	require.NoError(t, err)

	var notified bool
	var attempts int
	err = store.DB.QueryRowContext(ctx,
		"SELECT notified, notify_attempts FROM news WHERE id = $1", id,
	).Scan(&notified, &attempts)
	require.NoError(t, err)
	require.Equal(t, maxNotifyAttempts, attempts)
	require.True(t, notified, "item should be marked as notified at max attempts")
}

func TestStore_IncrementNotifyAttempts_DoesNotMarkNotifiedBelowMax(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()
	id := insertTestNews(t, store, "Test Band - Below Max (2026)")

	err := store.IncrementNotifyAttempts(ctx, id)
	require.NoError(t, err)

	var notified bool
	err = store.DB.QueryRowContext(ctx,
		"SELECT notified FROM news WHERE id = $1", id,
	).Scan(&notified)
	require.NoError(t, err)
	require.False(t, notified, "item should NOT be marked as notified below max attempts")
}

func TestStore_GetWithNotifyFlag_ExcludesMaxAttempts(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()

	// insert item at max attempts
	id := insertTestNews(t, store, "Test Band - Exhausted (2026)")
	_, err := store.DB.ExecContext(ctx,
		"UPDATE news SET notify_attempts = $1 WHERE id = $2", maxNotifyAttempts, id)
	require.NoError(t, err)

	// insert item with 0 attempts
	insertTestNews(t, store, "Test Band - Fresh (2026)")

	items, err := store.GetWithNotifyFlag(ctx)
	require.NoError(t, err)
	require.Len(t, items, 1)
	require.Equal(t, "Test Band - Fresh (2026)", items[0].Title)
}

func TestStore_GetWithNotifyFlag_RespectsRetryTime(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()

	id := insertTestNews(t, store, "Test Band - Waiting (2026)")

	// set next retry to future
	_, err := store.DB.ExecContext(ctx,
		"UPDATE news SET notify_attempts = 1, notify_next_retry = datetime('now', '+1 hours') WHERE id = $1", id)
	require.NoError(t, err)

	items, err := store.GetWithNotifyFlag(ctx)
	require.NoError(t, err)
	require.Empty(t, items, "item with future retry time should be excluded")
}
