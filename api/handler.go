package api

import (
	"encoding/json"
	"net/http"
)

type LinkStorage interface {
	Get(code string) (string, error)
	Put(code, url string) error
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

type CreateLinkRequest struct {
	URL  string `json:"url"`
	Code string `json:"code"`
}

func CreateLinkHandler(storage LinkStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body := r.Body
		defer body.Close()
		var requestBody CreateLinkRequest
		if err := json.NewDecoder(body).Decode(&requestBody); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		storage.Put(requestBody.Code, requestBody.URL)
		w.WriteHeader(http.StatusCreated)
	}
}
