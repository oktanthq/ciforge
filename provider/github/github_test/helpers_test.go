package github_test

import (
	"archive/zip"
	"bytes"
	"net/http"
	"testing"
)

func assertPath(t *testing.T, r *http.Request, expected string) {
	t.Helper()
	if r.URL.Path != expected {
		t.Errorf("path = %q, want %q", r.URL.Path, expected)
	}
}

func assertMethod(t *testing.T, r *http.Request, expected string) {
	t.Helper()
	if r.Method != expected {
		t.Errorf("method = %q, want %q", r.Method, expected)
	}
}

func assertAuth(t *testing.T, r *http.Request) {
	t.Helper()
	if r.Header.Get("Authorization") == "" {
		t.Error("missing Authorization header")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func createTestZip(files map[string]string) []byte {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	for name, content := range files {
		f, err := w.Create(name)
		if err != nil {
			panic(err)
		}
		f.Write([]byte(content))
	}
	w.Close()
	return buf.Bytes()
}
