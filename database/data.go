package database

type MythologicalCreature struct {
	Name      string                  `json:"name" validate:"required"`
	Nation    string                  `json:"nation" validate:"required"`
	Power     string                  `json:"power"`
	Type      string                  `json:"type" validate:"required"`
	Relations []*MythologicalCreature `json:"relations"`
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
