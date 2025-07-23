package models

import (
	"time"
	"database/sql"
)

type URL struct {
	ID int
	OriginalURL string
	ShortCode string
	UserID int
	CreatedAt time.Time
}

func (u *URL) CreateShortURL(db *sql.DB) error {
	query := `INSERT INTO url(original_url, short_code, user_id) VALUES($1, $2, $3) RETURNING id`
	return db.QueryRow(query, u.OriginalURL, u.ShortCode, u.UserID).Scan(&u.ID)
}

func (u *URL) GetOriginalURL(db *sql.DB) error {
	query := `SELECT original_url from url WHERE short_code = $1`
	return db.QueryRow(query, u.ShortCode).Scan(&u.OriginalURL)
}