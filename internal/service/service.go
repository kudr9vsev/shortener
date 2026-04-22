package service

import (
	"errors"
	"net/url"
	"strings"

	"github.com/kudr9vsev/shortener/internal/model"
)

type LinkService struct {
	repo MemoryLinkRepository
}

func NewLinkService(repo MemoryLinkRepository) *LinkService {
	return &LinkService{
		repo: repo,
	}
}

type MemoryLinkRepository interface {
	Create(link *model.Link) error
	GetByHash(string) (string, error)
}

func (s *LinkService) CreateLink(url string) (*model.Link, error) {
	if url == "" {
		return nil, errors.New("url can be empty")
	}

	if !IsURL(url) {
		return nil, errors.New("invalid url")
	}

	link := model.NewLink(url)

	if err := s.repo.Create(link); err != nil {
		return nil, err

	}

	return link, nil
}

func (s *LinkService) GetOriginalURL(hash string) (string, error) {
	if hash == "" {
		return "", errors.New("hash can be empty")
	}

	url, err := s.repo.GetByHash(hash)

	if err != nil {
		return "", err
	}
	return url, nil
}

func IsURL(rawUrl string) bool {
	if rawUrl == "" {
		return false
	}

	u, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	host := u.Hostname()
	if host == "" {
		return false
	}

	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		return false
	}

	for _, p := range parts {
		if p == "" {
			return false
		}
	}

	tld := parts[len(parts)-1]
	if len(tld) < 2 {
		return false
	}

	return true
}
