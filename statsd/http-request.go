package statsd

import (
	"net/http"
	"time"
)

type HttpHandlerStats struct {
	Handler http.Handler
}

func (s HttpHandlerStats) IsTimed(stat string) HttpHandlerStats {
	_next := s.Handler
	s.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		starttime := time.Now()
		_next.ServeHTTP(w, r)
		elapsedtime := time.Since(starttime)
		println(elapsedtime)
		Client.Timing("", int64(elapsedtime/time.Millisecond))
	})
	return s
}

func (s HttpHandlerStats) Done() http.Handler{
	return s.Handler
}

func (s HttpHandlerStats) IsCounted(stat string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Client != nil {
			Client.Increment(stat)
		}
		s.Handler.ServeHTTP(w, r)
	})
}

func Timed(stat string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		starttime := time.Now()
		next.ServeHTTP(w, r)
		elapsedtime := time.Since(starttime)
		println(elapsedtime)
		Client.Timing("", int64(elapsedtime/time.Millisecond))
	})
}

func CounterHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Client != nil {
			Client.Increment("")
		}
	})
}

func TimingHandler(stats string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		starttime := time.Now()
		next.ServeHTTP(w, r)
		elapsedtime := time.Since(starttime)
		println(elapsedtime)
		Client.Timing("", int64(elapsedtime/time.Millisecond))
	})
}
