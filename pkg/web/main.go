package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/sevilha/coffee/pkg/web/router"
)

func main() {
	r := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	router.MakeCoffeeHandler(r, n)
	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
