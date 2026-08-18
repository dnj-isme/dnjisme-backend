package external

// from SWGOH_FETCH_PLAYER_URL
type PlayerData struct {
	Mods struct {
		Profiles []struct {
			AllyCode int    `json:"allyCode"`
			Name     string `json:"name"`
			Guild    string `json:"guild"`
			Mods     []struct {
				SecondaryType1    string `json:"secondaryType_1"`
				SecondaryValue1   string `json:"secondaryValue_1"`
				SecondaryRoll1    string `json:"secondaryRoll_1"`
				SecondaryType2    string `json:"secondaryType_2"`
				SecondaryValue2   string `json:"secondaryValue_2"`
				SecondaryRoll2    string `json:"secondaryRoll_2"`
				SecondaryType3    string `json:"secondaryType_3"`
				SecondaryValue3   string `json:"secondaryValue_3"`
				SecondaryRoll3    string `json:"secondaryRoll_3"`
				SecondaryType4    string `json:"secondaryType_4"`
				SecondaryValue4   string `json:"secondaryValue_4"`
				SecondaryRoll4    string `json:"secondaryRoll_4"`
				PrimaryBonusType  string `json:"primaryBonusType"`
				PrimaryBonusValue string `json:"primaryBonusValue"`
				ModUID            string `json:"mod_uid"`
				Slot              string `json:"slot"`
				Set               string `json:"set"`
				Level             int    `json:"level"`
				Pips              int    `json:"pips"`
				CharacterID       string `json:"characterID"`
				Tier              int    `json:"tier"`
				ReRolledCount     int    `json:"reRolledCount"`
			} `json:"mods"`
			Characters []struct {
				ID        string        `json:"id"`
				BaseID    string        `json:"baseId"`
				Level     int           `json:"level"`
				Rarity    int           `json:"rarity"`
				GearLevel int           `json:"gearLevel"`
				Power     int           `json:"power"`
				RelicTier int           `json:"relicTier"`
				Equipment []interface{} `json:"equipment"`
				Stats     struct {
					Base struct {
						Health                    int     `json:"health"`
						Protection                int     `json:"protection"`
						Speed                     int     `json:"speed"`
						Potency                   float64 `json:"potency"`
						Tenacity                  float64 `json:"tenacity"`
						PhysicalDamage            int     `json:"Physical Damage"`
						PhysicalCriticalChance    int     `json:"Physical Critical Chance"`
						Armor                     int     `json:"armor"`
						SpecialDamage             int     `json:"Special Damage"`
						SpecialCriticalChance     int     `json:"Special Critical Chance"`
						Resistance                int     `json:"resistance"`
						CriticalDamage            float64 `json:"Critical Damage"`
						PhysicalCriticalAvoidance int     `json:"Physical Critical Avoidance"`
						PhysicalAccuracy          int     `json:"Physical Accuracy"`
					} `json:"base"`
				} `json:"stats"`
			} `json:"characters"`
		} `json:"profiles"`
	} `json:"mods"`
	ResponseMessage string `json:"responseMessage"`
	ResponseCode    int    `json:"responseCode"`
	ErrorMessage    string `json:"errorMessage"`
	ErrorSeverity   int    `json:"errorSeverity"`
}

// from SWGOH_FETCH_CHARACTER_URL
type CharacterData struct {
	Units []struct {
		BaseID           string        `json:"baseId"`
		Name             string        `json:"name"`
		BaseImage        string        `json:"baseImage"`
		CombatType       int           `json:"combatType"`
		Alignment        int           `json:"alignment"`
		GalacticLegend   bool          `json:"galacticLegend"`
		FleetCommander   bool          `json:"fleetCommander"`
		HasLeaderAbility bool          `json:"hasLeaderAbility"`
		HasZetaLead      bool          `json:"hasZetaLead"`
		HasOmicronLead   bool          `json:"hasOmicronLead"`
		GgImage          interface{}   `json:"ggImage"`
		ShipBaseID       interface{}   `json:"shipBaseId"`
		ShipSlot         int           `json:"shipSlot"`
		Description      string        `json:"description"`
		IsEraCharacter   bool          `json:"isEraCharacter"`
		Abbreviations    []string      `json:"abbreviations"`
		Affiliation      []interface{} `json:"affiliation"`
		Profession       []struct {
			Key     string `json:"key"`
			Display string `json:"display"`
			Leader  bool   `json:"leader"`
			Mode    int    `json:"mode"`
		} `json:"profession"`
		Role []struct {
			Key     string `json:"key"`
			Display string `json:"display"`
			Leader  bool   `json:"leader"`
			Mode    int    `json:"mode"`
		} `json:"role"`
		Species []struct {
			Key     string `json:"key"`
			Display string `json:"display"`
			Leader  bool   `json:"leader"`
			Mode    int    `json:"mode"`
		} `json:"species"`
		Other interface{} `json:"other"`
		Zeta  []struct {
			Key     string `json:"key"`
			Display string `json:"display"`
			Leader  bool   `json:"leader"`
			Mode    int    `json:"mode"`
		} `json:"zeta"`
		Omicron []struct {
			Key     string `json:"key"`
			Display string `json:"display"`
			Leader  bool   `json:"leader"`
			Mode    int    `json:"mode"`
		} `json:"omicron"`
	} `json:"units"`
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ErrorMessage    string `json:"errorMessage"`
	ErrorSeverity   int    `json:"errorSeverity"`
}
