package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestYAMLHandler(t *testing.T) {
	yml, err := os.ReadFile("testdata/sample.yaml")
	if err != nil {
		t.Error("unexpected error:", err)
	}

	handler, err := YAMLHandler(yml, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Fallback"))
	}))
	if err != nil {
		t.Error("unexpected error:", err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/goo", nil)
	handler(w, req)

	resp := w.Result()
	if code := resp.StatusCode; code != http.StatusFound {
		t.Errorf("expected %d, get %d", http.StatusFound, code)
	}

	expectedURL := "https://google.com"
	respURL, err := resp.Location()
	if err != nil {
		t.Errorf("no redirect!")
	} else if respURL.String() != expectedURL {
		t.Errorf(`expected %v, get %v`, expectedURL, respURL)
	}
}

func TestJSONHandler(t *testing.T) {
	yml, err := os.ReadFile("testdata/sample.json")
	if err != nil {
		t.Error("unexpected error:", err)
	}

	handler, err := JSONHandler(yml, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Fallback"))
	}))
	if err != nil {
		t.Error("unexpected error:", err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/goo", nil)
	handler(w, req)

	resp := w.Result()
	if code := resp.StatusCode; code != http.StatusFound {
		t.Errorf("expected %d, get %d", http.StatusFound, code)
	}

	expectedURL := "https://google.com"
	respURL, err := resp.Location()
	if err != nil {
		t.Errorf("no redirect!")
	} else if respURL.String() != expectedURL {
		t.Errorf(`expected %v, get %v`, expectedURL, respURL)
	}
}
