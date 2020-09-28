package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vSterlin/mythology/database"
)

func GetHanlder(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(database.Database)
}
