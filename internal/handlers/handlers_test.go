package handlers

import (
	"bytes"
	store "devops/internal/storage"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestMetricsResourceUpdate_Update(t *testing.T) {
	m := store.New()
	m.Init()
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		url  string
		want want
	}{
		{
			name: "positive test #1",
			url:  fmt.Sprintf("/update/counter/PollCount/%d", 10),
			want: want{
				code:        200,
				response:    `{"status":"ok"}`,
				contentType: "text/plain",
			},
		},
		{
			name: "positive test #2",
			url:  fmt.Sprintf("/update/gauge/Alloc/%g", 10.1),
			want: want{
				code:        200,
				response:    `{"status":"ok"}`,
				contentType: "text/plain",
			},
		},
		{
			name: "negative test #1",
			url:  "/update/counter/testCounter/none",
			want: want{
				code:        400,
				response:    `{"status":"Bad request"}`,
				contentType: "text/plain",
			},
		},
		{
			name: "negative test #2",
			url:  "/update/counter/",
			want: want{
				code:        404,
				response:    `{"status":"Not found"}`,
				contentType: "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.url, bytes.NewBufferString(""))
			w := httptest.NewRecorder()
			// определяем хендлер
			h := chi.NewRouter()
			h.Mount("/update", MetricsResourceUpdate{M: m}.Routes())
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}
		})
	}
}

func TestMetricsResourceUpdate_Value(t *testing.T) {
	m := store.New()
	m.Init()
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		url  string
		want want
	}{
		{
			name: "negative test #1",
			url:  "/value/counter/testSetGet33",
			want: want{
				code:        404,
				response:    `{"status":"ok"}`,
				contentType: "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.url, bytes.NewBufferString(""))
			w := httptest.NewRecorder()
			// определяем хендлер
			h := chi.NewRouter()
			h.Mount("/value", MetricsResource{M: m}.Routes())
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			if res.StatusCode != tt.want.code {
				t.Log(w.Body)
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}
		})
	}
}
