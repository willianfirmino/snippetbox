package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	code, _, body := ts.get(t, "/ping")
	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}
	if body != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestHome(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	t.Run("home", func(t *testing.T) {
		code, _, body := ts.get(t, "/")
		if code != http.StatusOK {
			t.Errorf("want %d; got %d", http.StatusOK, code)
		}
		if !bytes.Contains([]byte(body), []byte("An old silent pond")) {
			t.Errorf("want body to contain %q", "An old silent pond")
		}
	})
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}
			if !bytes.Contains([]byte(body), []byte(tt.wantBody)) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}
}

func TestSignup(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/user/signup")
	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}
	if !bytes.Contains([]byte(body), []byte("Signup")) {
		t.Errorf("want body to contain %q, but got %q", "Signup", body)
	}
}
