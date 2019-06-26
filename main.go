package main

import (
	"net/http"
	"time"

	"github.com/muhoro/statsdtest/statsd"
)

type StatsDClient struct{}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", statsd.HttpHandlerStats{Handler: http.HandlerFunc(test)}.IsTimed("").Done())
	http.ListenAndServe(":1234", mux)
}

func test(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2000)
	w.Write([]byte("Tuko sawa"))
}
