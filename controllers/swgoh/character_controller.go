package swgoh

import (
	"dnj-backend/config"
	"dnj-backend/models"
	"dnj-backend/models/external"
	"dnj-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchPlayerData(c *gin.Context) {
	allycode := c.Param("allycode")

	if allycode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Allycode is required",
		})
		return
	}

	endpoint := config.LoadEnv().SWGoH.FetchPlayerURL
	payload := external.FetchPlayerDto{
		Action: "getprofile",
		Payload: external.FetchPlayerPayload{
			AllyCode: allycode,
		},
	}

	var res external.PlayerData

	code, err := utils.FetchAPI(endpoint, &utils.RequestOptions{
		Method: "POST",
		Body:   payload,
	}, &res)

	if err != nil || code != http.StatusOK || res.ErrorMessage == "ERROR" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch player data",
		})
		return
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
			response.CharacterIds = append(response.CharacterIds, character.BaseID)
			switch character.GearLevel {
			case 13:
				response.RelicCharacterIds = append(response.RelicCharacterIds, character.BaseID)
			case 12:
				response.Gear12CharacterIds = append(response.Gear12CharacterIds, character.BaseID)
			default:
				response.NonRelicCharacterIds = append(response.NonRelicCharacterIds, character.BaseID)
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

func FetchAllCharacters(c *gin.Context) {
	endpoint := config.LoadEnv().SWGoH.FetchCharacterURL

	var res external.CharacterData

	code, err := utils.FetchAPI(endpoint, &utils.RequestOptions{
		Method: "GET",
	}, &res)

	if err != nil || code != http.StatusOK || res.ErrorMessage == "ERROR" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch character data",
		})
		return
	}

	var response []string
	for _, character := range res.Units {
		response = append(response, character.BaseID)
	}
	c.JSON(http.StatusOK, response)
}
