package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/vSterlin/mythology/database"

	"github.com/gorilla/mux"
)

func GetOneHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	for _, item := range database.Database {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			json.NewEncoder(w).Encode(item)

		}
	}

}
