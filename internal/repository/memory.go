package repository

import (
	"errors"

	"github.com/kudr9vsev/shortener/internal/model"
)

type MemoryLinkRepository struct {
	data map[string]string
}

func NewMemoryLinkRepository() *MemoryLinkRepository {
	return &MemoryLinkRepository{
		data: make(map[string]string),
	}
}

func (r *MemoryLinkRepository) Create(link *model.Link) error {
	r.data[link.Hash] = link.URL
	return nil
}

func (r *MemoryLinkRepository) GetByHash(hash string) (string, error) {
	url, ok := r.data[hash]
	if !ok {
		return "", errors.New("link not found")
	}

	return url, nil
}
