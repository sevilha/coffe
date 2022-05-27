package handlers

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/sevilha/coffee/pkg/model/coffee"
)

func MakeCoffeeHandler(r *mux.Router, n *negroni.Negroni, service coffee.UseCase) {
	//r.Use(commonMiddleware)
	r.HandleFunc("/v1/coffee", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		negroni.Wrap(getAllCoffees(service))
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/v1/coffee/{id}", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		negroni.Wrap(getCoffee(service))
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/v1/coffee", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		negroni.Wrap(storeCoffee(service))
	}).Methods("POST", "OPTIONS")

	r.HandleFunc("/v1/coffee/{id}", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		negroni.Wrap(updateCoffee(service))
	}).Methods("PUT", "OPTIONS")

	r.HandleFunc("/v1/coffee/{id}", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		negroni.Wrap(removeCoffee(service))
	}).Methods("DELETE", "OPTIONS")
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getAllCoffees(service coffee.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
func getCoffee(service coffee.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
func storeCoffee(service coffee.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
func updateCoffee(service coffee.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
func removeCoffee(service coffee.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
