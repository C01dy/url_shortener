package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockStorage struct {
	PutFunc func(code, url string) error
	GetFunc func(code string) (string, error)
}

func (m *MockStorage) Put(code, url string) error {
	return m.PutFunc(code, url)
}

func (m *MockStorage) Get(code string) (string, error) {
	return m.GetFunc(code)
}

func TestCreateLinkHandler(t *testing.T) {
	t.Run("should create link successfully", func(t *testing.T) {
		mockStorage := &MockStorage{
			PutFunc: func(code, url string) error {
				if url != "https://google.com" {
					t.Errorf("expected url to be 'https://google.com', got %s", url)
				}
				return nil
			},
		}

		handler := CreateLinkHandler(mockStorage)

		requestBody := strings.NewReader(`{"url":"https://google.com"}`)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/links", requestBody)
		rr := httptest.NewRecorder()

		handler(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		var resp CreateLinkResponse
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("could not decode response json: %v", err)
		}

		if resp.ShortURL == "" {
			t.Error("short_url should not be empty")
		}
		
	})

	t.Run("should return bad request for invalid json", func(t *testing.T) {
		mockStorage := &MockStorage{}

		handler := CreateLinkHandler(mockStorage)

		requestBody := strings.NewReader(`{"url":}`)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/links", requestBody)
		rr := httptest.NewRecorder()

		handler(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		var resp CreateLinkResponse
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.Fatalf("could not decode response json: %v", err)
		}

		if resp.ShortURL == "" {
			t.Error("short_url should not be empty")
		}
	})
}