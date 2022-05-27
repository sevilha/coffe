package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/sevilha/coffee/pkg/model/coffee"
	handlers "github.com/sevilha/coffee/pkg/web/handlers"
)

type controller struct {
	logger        *log.Logger
	nextRequestID func() string
	healthy       int64
}

type middleware func(http.Handler) http.Handler
type middlewares []middleware

func main() {
	db, err := sql.Open("sqlite3", "db/coffee-go.db")

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Server is starting...")

	c := &controller{logger: logger, nextRequestID: func() string { return strconv.FormatInt(time.Now().UnixNano(), 36) }}

	service := coffee.NewService(db)
	handlers.MakeCoffeeHandler(r, n, service)
	http.Handle("/", r)

	addr := ":4000"
	server := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         addr,
		Handler:      (middlewares{c.tracing, c.logging}).apply(r),
		ErrorLog:     logger,
	}

	logger.Printf("Server is ready to handle requests at %q\n", addr)
	atomic.StoreInt64(&c.healthy, time.Now().UnixNano())

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func (c *controller) logging(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func(start time.Time) {
			requestID := w.Header().Get("X-Request-Id")
			if requestID == "" {
				requestID = "unknown"
			}
			c.logger.Println(requestID, req.Method, req.URL.Path, req.RemoteAddr, time.Since(start))
		}(time.Now())
		hdlr.ServeHTTP(w, req)
	})
}

func (c *controller) tracing(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestID := req.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = c.nextRequestID()
		}
		w.Header().Set("X-Request-Id", requestID)
		hdlr.ServeHTTP(w, req)
	})
}

func (mws middlewares) apply(hdlr http.Handler) http.Handler {
	if len(mws) == 0 {
		return hdlr
	}
	return mws[1:].apply(mws[0](hdlr))
}
