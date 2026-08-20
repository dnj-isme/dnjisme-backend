package utils

import (
	"dnj-backend/config"
	"dnj-backend/models"
	"dnj-backend/models/external"
	"net/http"
	"strconv"
)

func GetCharacters() (external.CharacterData, error) {

	endpoint := config.LoadEnv().SWGoH.FetchCharacterURL

	var res external.CharacterData

	code, err := FetchAPI(endpoint, &RequestOptions{
		Method: "GET",
	}, &res)

	if err != nil || code != http.StatusOK || res.ErrorMessage == "ERROR" {
		return external.CharacterData{}, err
	}

	return res, nil
}

func GetSimplifiedCharacters() ([]models.SimplifiedCharacterData, error) {
	characters, err := GetCharacters()
	if err != nil {
		return nil, err
	}
	var simplifiedCharacters []models.SimplifiedCharacterData

	for _, unit := range characters.Units {
		simplifiedCharacters = append(simplifiedCharacters, models.SimplifiedCharacterData{
			BaseID:    unit.BaseID,
			Name:      unit.Name,
			BaseImage: unit.BaseImage,
		})
	}

	return simplifiedCharacters, nil
}

func GetPlayerData(allycode string) (models.PlayerDataResponse, error) {
	characterList, err := GetSimplifiedCharacters()
	if err != nil {
		return models.PlayerDataResponse{}, err
	}

	var res external.PlayerData

	endpoint := config.LoadEnv().SWGoH.FetchPlayerURL
	payload := external.FetchPlayerDto{
		Action: "getprofile",
		Payload: external.FetchPlayerPayload{
			AllyCode: allycode,
		},
	}

	code, err := FetchAPI(endpoint, &RequestOptions{
		Method: "POST",
		Body:   payload,
	}, &res)

	if err != nil || code != http.StatusOK || res.ErrorMessage == "ERROR" {
		return models.PlayerDataResponse{}, err
	}

	var response models.PlayerDataResponse

	for _, profile := range res.Mods.Profiles {
		if strconv.Itoa(profile.AllyCode) != allycode {
			continue
		}
		response.Allycode = strconv.Itoa(profile.AllyCode)
		response.Username = profile.Name
		response.GuildName = profile.Guild

		for _, character := range profile.Characters {

			var c models.SimplifiedCharacterData
			found := false
			
			for _, char := range characterList {
				if char.BaseID == character.BaseID {
					c = char
					found = true
					break
				}
			}

			if !found {
				c = models.SimplifiedCharacterData{
					BaseID:    character.BaseID,
					Name:      character.BaseID,
					BaseImage: "",
				}
			}

			response.Characters = append(response.Characters, c)
			switch character.GearLevel {
			case 13:
				response.RelicCharacters = append(response.RelicCharacters, c)
			case 12:
				response.Gear12Characters = append(response.Gear12Characters, c)
			default:
				response.NonRelicCharacters = append(response.NonRelicCharacters, c)
			}
		}
	}

	return response, nil
}
