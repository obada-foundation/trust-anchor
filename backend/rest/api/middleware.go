package api

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

// AppInfo adds custom app-info to the response header
func AppInfo(app, author, version string) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Author", author)
			w.Header().Set("App-Name", app)
			w.Header().Set("App-Version", version)
			if mhost := os.Getenv("MHOST"); mhost != "" {
				w.Header().Set("Host", mhost)
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}

// Ping middleware response with pong to /ping. Stops chain if ping request detected
func Ping(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" && strings.HasSuffix(strings.ToLower(r.URL.Path), "/ping") {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("pong"))
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// Recoverer is a middleware that recovers from panics, logs the panic and returns a HTTP 500 status if possible.
func Recoverer(l *log.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					l.Printf("request panic for %s from %s, %v", r.URL.String(), r.RemoteAddr, rvr)
					if rvr != http.ErrAbortHandler {
						l.Printf(string(debug.Stack()))
					}
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
