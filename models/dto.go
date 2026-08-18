package models

type RegisterUserDto struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginUserDto struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type PlayerDataResponse struct {
	Allycode             string   `json:"allycode"`
	Username             string   `json:"username"`
	GuildName            string   `json:"guildName"`
	RelicCharacterIds    []string `json:"relicCharacterIds"`
	Gear12CharacterIds   []string `json:"gear12CharacterIds"`
	NonRelicCharacterIds []string `json:"nonRelicCharacterIds"`
	CharacterIds         []string `json:"characterIds"`
}

type SaveTemplateDto struct {
	TemplateName string
	TemplateType string
	CreatedBy    uint
	Squads       []struct {
		SquadName      string
		SquadWeightGAC float64
		SquadWeightTW  float64
		Units          []struct {
			CharacterID    string
			UnitWeight     float64
			UnitOrder      int
			PriorityNumber int
		}
	}
}

type DeleteTemplateDto struct {
	TemplateName string
}
