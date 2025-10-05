package api

import (
	"math/rand"
	"encoding/json"
	"net/http"
	"time"
)

type LinkStorage interface {
	Get(code string) (string, error)
	Put(code, url string) error
}

type CreateLinkRequest struct {
	URL string `json:"url"`
}

type CreateLinkResponse struct {
	ShortURL string `json:"short_url"`
}

const codeAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 6

func generateShortCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = codeAlphabet[r.Intn(len(codeAlphabet))]
	}
	return string(b)
}

func RedirectHandler(storage LinkStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[1:]
		url, err := storage.Get(code)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	}
}

func CreateLinkHandler(storage LinkStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		defer body.Close()
		var requestBody CreateLinkRequest
		if err := json.NewDecoder(body).Decode(&requestBody); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if requestBody.URL == "" {
			RespondWithError(w, http.StatusBadRequest, "URL is required")
			return
		}
		code := generateShortCode()
		if err := storage.Put(code, requestBody.URL); err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not create short link")
			return
		}

		response := CreateLinkResponse{
			ShortURL: r.Host + "/" + code,
		}

		RespondWithJSON(w, http.StatusCreated, response)
	}
}
