package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/kudr9vsev/shortener/internal/model"
)

type LinkService interface {
	CreateLink(string) (*model.Link, error)
	GetOriginalURL(string) (string, error)
}

type LinkHandler struct {
	linkService LinkService
}

func NewLinkHandler(linkService LinkService) *LinkHandler {
	return &LinkHandler{
		linkService: linkService,
	}
}

func (h *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-type")

	if contentType == "" || contentType != "text/plain" {
		http.Error(w, "invalid Content-Type: expected text/plain", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "text/plain")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	normalizedBody := strings.TrimSpace(string(body))

	link, err := h.linkService.CreateLink(normalizedBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortUrl := "http://" + r.Host + "/" + link.Hash

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(shortUrl))

}

func (h *LinkHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/")
	if hash == "" {
		http.Error(w, "empty hash", http.StatusBadRequest)
		return
	}

	url, err := h.linkService.GetOriginalURL(hash)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
