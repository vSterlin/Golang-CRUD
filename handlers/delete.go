package handlers

import (
	"net/http"
	"strings"

	"github.com/vSterlin/mythology/database"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	for i, item := range database.Database {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			newSlice := database.Database[0:i]
			newSlice = append(newSlice, database.Database[i+1:]...)
			database.Database = newSlice
		}
	}

}
