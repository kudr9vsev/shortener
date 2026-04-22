package model

import (
	"math/rand"
	"time"
)

type Link struct {
	ID        int64     `db:"id"`
	URL       string    `db:"url"`
	Hash      string    `hash:"hash"`
	CreatedAt time.Time `db:"created_at"`
}

func NewLink(url string) *Link {
	return &Link{
		URL:       url,
		Hash:      RandString(5),
		CreatedAt: time.Now(),
	}
}

var letters = []rune("abcdefghijklmnopstqwxyzABCDEFGHIJKLMNOPRSTQWXYZ")

func RandString(n int) string {
	s := make([]rune, n)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
