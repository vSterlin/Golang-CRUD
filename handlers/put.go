package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vSterlin/mythology/database"
)

func PutHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	var changedCreature = new(database.MythologicalCreature)
	json.NewDecoder(r.Body).Decode(&changedCreature)

	for i, item := range database.Database {
		if strings.ToLower(item.Name) == strings.ToLower(name) {
			database.Database[i] = changedCreature
			json.NewEncoder(w).Encode(item)

		}
	}
}
