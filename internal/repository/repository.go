package repository

import (
	"database/sql"
	"errors"

	"github.com/kudr9vsev/shortener/internal/model"
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (r *LinkRepository) Create(link *model.Link) error {
	query := `INSERT INTO links (url, hash, created_at) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(
		query,
		link.URL,
		link.Hash,
		link.CreatedAt,
	).Scan(&link.ID)
	return err
}

func (r *LinkRepository) GetByHash(hash string) (string, error) {
	query := `SELECT url FROM links WHERE hash = $1`

	var OriginalUrl string

	err := r.db.QueryRow(query, hash).Scan(&OriginalUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("link not found")
		}
		return "", err
	}

	return OriginalUrl, nil
}
