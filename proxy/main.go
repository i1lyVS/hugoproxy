package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"


	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	proxy := NewReverseProxy("hugo", "1313")

	r.Get("/api/", Handler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		proxy.ReverseProxy(nil).ServeHTTP(w, r)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server startup error: %v", err)
	}

	go WorkerTest()
}

const content = ``

func WorkerTest() {
	t := time.NewTicker(1 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprint(content, b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
