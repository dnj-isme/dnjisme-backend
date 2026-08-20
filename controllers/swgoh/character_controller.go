package swgoh

import (
	"dnj-backend/utils"
	"net/http"

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

	response, err := utils.GetPlayerData(allycode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch player data",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func FetchAllCharacters(c *gin.Context) {
	var response, err = utils.GetSimplifiedCharacters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch character data",
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
