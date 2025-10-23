package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"url-shortener/model"
	"url-shortener/repository"
	"url-shortener/utils"

	"github.com/gorilla/mux"
)

type URLHandler struct {
	repo repository.PostgresRepository
}

func NewURLHandler() *URLHandler {
	repo, err := repository.NewPostgresRepository()
	if err != nil {
		fmt.Println("Error DB:", err.Error())
	}
	return &URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/list":
			h.ListUrls(w, r)
		case "/clicks":
			h.ListTopClicked(w, r)
		default:
			h.Redirect(w, r)
		}
	case http.MethodPost:
		h.ShortenURL(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortcode := params["short"]
	url, err := h.repo.LoadURL(shortcode)
	if err != nil {
		http.Error(w, "Error loading URL: "+err.Error(), http.StatusNotFound)
	} else if time.Now().After(url.ExpiresAt) {
		http.Error(w, "URL expired", http.StatusGone)
		return
	}
	err = h.repo.IncrementClicks(shortcode)
	if err != nil {
		http.Error(w, "Error updating clicks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url.Link, http.StatusFound)
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var reqData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	urlLink, err := utils.PrepareURL(reqData["url"])
	if err != nil {
		http.Error(w, "Error validating URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	shortcode := utils.CreateShortcode([]byte(urlLink))

	expDuration, err := time.ParseDuration(reqData["expires_in"])
	if err != nil {
		http.Error(w, "Error parsing duration: "+err.Error(), http.StatusBadRequest)
		return
	}

	url := model.Url{
		Link:        urlLink,
		Hash:        shortcode,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(expDuration),
		ClickCounts: 0,
	}

	if err := h.repo.SaveURL(url); err != nil {
		http.Error(w, "Error saving URL: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *URLHandler) ListUrls(w http.ResponseWriter, r *http.Request) {
	urlMap, err := h.repo.LoadAll(false)
	if err != nil {
		http.Error(w, "Error loading urls: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(urlMap); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *URLHandler) ListTopClicked(w http.ResponseWriter, r *http.Request) {
	urlMap, err := h.repo.LoadAll(true)
	if err != nil {
		http.Error(w, "Error loading urls: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(urlMap); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
