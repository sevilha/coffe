package router

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeCoffeeHandler(r *mux.Router, n *negroni.Negroni) {
	r.HandleFunc("/v1/coffee", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/v1/coffee/{id}", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/v1/coffee", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("POST", "OPTIONS")

	r.HandleFunc("/v1/coffee/{id}", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("PUT", "OPTIONS")

	r.HandleFunc("/v1/coffee/{id}", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("DELETE", "OPTIONS")
}
