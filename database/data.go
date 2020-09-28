package database

import "github.com/Kamva/mgm"

type MythologicalCreature struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string                  `json:"name" bson:"name" validate:"required"`
	Nation           string                  `json:"nation" bson:"nation" validate:"required"`
	Power            string                  `json:"power" bson:"power"`
	Type             string                  `json:"type" bson:"type" validate:"required"`
	Relations        []*MythologicalCreature `json:"relations" bson:"relations"`
}

var Database = []*MythologicalCreature{
	&MythologicalCreature{
		Name:   "Zeus",
		Nation: "Greek",
		Type:   "God",
		Power:  "Lightning",
	},
	&MythologicalCreature{
		Name:   "Perkunas",
		Nation: "Lithuanian",
		Type:   "God",
		Power:  "Lightning",
	},
}
