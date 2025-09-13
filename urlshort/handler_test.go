package urlshort

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func fallbackTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func TestMapHandler(t *testing.T) {
	pathsToUrls := map[string]string{
		"/gophercises": "https://gophercises.com/",
	}
	handler := MapHandler(pathsToUrls, http.HandlerFunc(fallbackTest))

	t.Run("should redirect when path is found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/gophercises", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
		}

		expectedURL := "https://gophercises.com/"
		if location := rr.Header().Get("Location"); location != expectedURL {
			t.Errorf("handler returned wrong status code: got %v want %v", location, expectedURL)
		}
	})

	t.Run("should call fallback when path is not found", func(t *testing.T) {
		notExistedPath := "/unknown"
		req := httptest.NewRequest(http.MethodGet, notExistedPath, nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)

		}
	})
}
