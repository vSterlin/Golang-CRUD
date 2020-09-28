package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vSterlin/mythology/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetHanlder).Methods(http.MethodGet)
	r.HandleFunc("/{name}", handlers.GetOneHandler).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.PostHandler).Methods(http.MethodPost)
	r.HandleFunc("/{name}", handlers.PutHandler).Methods(http.MethodPut)
	r.HandleFunc("/{name}", handlers.DeleteHandler).Methods(http.MethodDelete)
	http.ListenAndServe("localhost:8080", r)
}
