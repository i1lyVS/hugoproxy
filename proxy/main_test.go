package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestHandleHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/", nil)
	rr := httptest.NewRecorder()
	handleHello(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rr.Code)
	}

	expectedBody := "Hello from API"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rr.Body.String())
	}
}

func TestNewReverseProxy(t *testing.T) {
	type args struct {
		host string
		port string
	}
	tests := []struct {
		name string
		args args
		want *ReverseProxy
	}{
		{
			name: "base",
			args: args{
				host: "test",
				port: "322",
			},
			want: &ReverseProxy{
				host: "test",
				port: "322",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReverseProxy(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReverseProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseProxy_ReverseProxy(t *testing.T) {
	type fields struct {
		host string
		port string
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		requestURL string
		want       string
	}{
		{
			name: "Base case - API route",
			fields: fields{
				host: "localhost",
				port: "8080",
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Next handler response"))
				}),
			},
			requestURL: "/api/test",
			want:       "Next handler response",
		},
		{
			name: "Static file route - proxied",
			fields: fields{
				host: "localhost",
				port: "8080",
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("This should not be called"))
				}),
			},
			requestURL: "/static/file.css",
			want:       "Proxied response",
		},
		{
			name: "Matching host - next handler",
			fields: fields{
				host: "localhost",
				port: "8080",
			},
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Next handler response"))
				}),
			},
			requestURL: "http://localhost:8080/static/file.css",
			want:       "Next handler response",
		},
	}

	// Запускаем сервер-заглушку
	proxiedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Proxied response"))
	}))
	defer proxiedServer.Close()

	// Извлекаем хост и порт из URL заглушки
	proxiedURL, _ := url.Parse(proxiedServer.URL)
	host, port, _ := net.SplitHostPort(proxiedURL.Host)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Используем заглушку как проксируемый сервер
			if tt.name == "Static file route - proxied" || tt.name == "Invalid URL" {
				tt.fields.host = host
				tt.fields.port = port
			}

			rp := &ReverseProxy{
				host: tt.fields.host,
				port: tt.fields.port,
			}

			req := httptest.NewRequest("GET", tt.requestURL, nil)
			w := httptest.NewRecorder()

			rp.ReverseProxy(tt.args.next).ServeHTTP(w, req)

			got := w.Body.String()
			if got != tt.want {
				t.Errorf("ReverseProxy() = %v; want %v", got, tt.want)
			}
		})
	}
}
