package main

import (
	"fmt"
	"net/http"

	"github.com/Kamva/mgm"
	"github.com/gorilla/mux"
	"github.com/vSterlin/mythology/handlers"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	err := mgm.SetDefaultConfig(nil, "mythology", options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetHanlder).Methods(http.MethodGet)
	r.HandleFunc("/{name}", handlers.GetOneHandler).Methods(http.MethodGet)
	r.HandleFunc("/", handlers.PostHandler).Methods(http.MethodPost)
	r.HandleFunc("/{name}", handlers.PutHandler).Methods(http.MethodPut)
	r.HandleFunc("/{name}", handlers.DeleteHandler).Methods(http.MethodDelete)
	http.ListenAndServe("localhost:8080", r)
}
