package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

//Application is the main structure of this application
type Application struct {
	db     *Database
	mux    *chi.Mux
	port   string
	server *http.Server
}

//NewApplication creates new application
func NewApplication() *Application {
	var err error

	a := &Application{}
	a.mux = a.setUpMux()
	a.setUpRoute()

	a.db, err = NewDefaultDatabase()
	check(err)

	return a
}

//Start starts application
func (a *Application) Start() {
	a.port = env("PORT", "8080")
	a.server = &http.Server{
		Addr:         "0.0.0.0:" + a.port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.handleRequest(a.mux),
	}

	go func() {
		llog("info", "starting on port %s", a.port)
		if err := a.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				llog("error", "failed to start: %v", err)
				os.Exit(1)
			}
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	a.server.Shutdown(ctx)
	llog("info", "shutting down")
	time.Sleep(50 * time.Millisecond)
	os.Exit(0)
}

func (a *Application) handleRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS & other headers
		w.Header().Set("Service-Worker-Allowed", "/")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if (*r).Method == "OPTIONS" {
			return
		}
		start := time.Now().UnixNano()

		handler.ServeHTTP(w, r)
		partnershipKey := r.Header.Get("X-Partnership-Key")
		ms := (time.Now().UnixNano() - start) / int64(time.Millisecond)
		llog("info", "request %s %s %s (%dms)", r.Method, r.URL.Path, partnershipKey, ms)
	})
}
