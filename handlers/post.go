package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/vSterlin/mythology/database"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var newCreature = new(database.MythologicalCreature)
	json.NewDecoder(r.Body).Decode(&newCreature)
	validate := validator.New()
	err := validate.Struct(newCreature)
	if err != nil {
		errorMessage := fmt.Sprintf(err.Error())
		w.Write([]byte(errorMessage))
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field() + " is " + err.Tag())
		}
	}
	database.Database = append(database.Database, newCreature)
}
