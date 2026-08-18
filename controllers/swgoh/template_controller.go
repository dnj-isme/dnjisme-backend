package swgoh

import (
	"errors"

	"dnj-backend/database"
	"dnj-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveTemplate(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	var dto models.SaveTemplateDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// A. SquadTemplate (upsert)
	var template models.SquadTemplate
	err := database.DB.Where("created_by = ? AND template_name = ?", user.ID, dto.TemplateName).First(&template).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		template = models.SquadTemplate{
			CreatedBy:    user.ID,
			TemplateName: dto.TemplateName,
			TemplateType: dto.TemplateType,
		}
		if err := database.DB.Create(&template).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create template"})
			return
		}
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch template"})
		return
	} else if err := database.DB.Model(&template).Updates(map[string]interface{}{
		"template_type": dto.TemplateType,
	}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update template"})
		return
	}

	// Delete stale squads
	squadNames := make([]string, 0, len(dto.Squads))
	for _, squad := range dto.Squads {
		squadNames = append(squadNames, squad.SquadName)
	}

	var staleSquads []models.Squad
	staleSquadQuery := database.DB.Where("template_id = ?", template.ID)
	if len(squadNames) > 0 {
		staleSquadQuery = staleSquadQuery.Where("squad_name NOT IN ?", squadNames)
	}
	if err := staleSquadQuery.Find(&staleSquads).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch stale squads"})
		return
	}
	for _, staleSquad := range staleSquads {
		if err := database.DB.Where("squad_id = ?", staleSquad.ID).Delete(&models.SquadUnit{}).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete stale squad units"})
			return
		}
		if err := database.DB.Delete(&staleSquad).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete stale squad"})
			return
		}
	}

	// B. Squads from DTO (upsert)
	for _, squad := range dto.Squads {
		// Squad (Upsert)
		var existingSquad models.Squad
		err := database.DB.Where("template_id = ? AND squad_name = ?", template.ID, squad.SquadName).First(&existingSquad).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			existingSquad = models.Squad{
				TemplateID:     template.ID,
				SquadName:      squad.SquadName,
				SquadWeightGAC: squad.SquadWeightGAC,
				SquadWeightTW:  squad.SquadWeightTW,
			}
			if err := database.DB.Create(&existingSquad).Error; err != nil {
				c.JSON(500, gin.H{"error": "Failed to create squad"})
				return
			}
		} else if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch squad"})
			return
		} else if err := database.DB.Model(&existingSquad).Updates(map[string]any{
			"squad_weight_gac": squad.SquadWeightGAC,
			"squad_weight_tw":  squad.SquadWeightTW,
		}).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update squad"})
			return
		}

		// Delete stale SquadUnits
		var staleUnits []models.SquadUnit
		staleUnitQuery := database.DB.Where("squad_id = ?", existingSquad.ID)
		if len(squad.Units) > 0 {
			characterIDs := make([]string, 0, len(squad.Units))
			for _, unit := range squad.Units {
				characterIDs = append(characterIDs, unit.CharacterID)
			}
			staleUnitQuery = staleUnitQuery.Where("character_id NOT IN ?", characterIDs)
		}
		if err := staleUnitQuery.Find(&staleUnits).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch stale squad units"})
			return
		}
		for _, staleUnit := range staleUnits {
			if err := database.DB.Delete(&staleUnit).Error; err != nil {
				c.JSON(500, gin.H{"error": "Failed to delete stale squad unit"})
				return
			}
		}

		// C. SquadUnits from DTO
		for _, unit := range squad.Units {

			// SquadUnit (Upsert)
			var existingUnit models.SquadUnit
			err := database.DB.Where("squad_id = ? AND character_id = ?", existingSquad.ID, unit.CharacterID).First(&existingUnit).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newUnit := models.SquadUnit{
					SquadID:        existingSquad.ID,
					CharacterID:    unit.CharacterID,
					UnitWeight:     unit.UnitWeight,
					UnitOrder:      unit.UnitOrder,
					PriorityNumber: unit.PriorityNumber,
				}
				if err := database.DB.Create(&newUnit).Error; err != nil {
					c.JSON(500, gin.H{"error": "Failed to create squad unit"})
					return
				}
			} else if err != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch squad unit"})
				return
			}
		}
	}
	c.JSON(200, gin.H{"status": "success", "template_id": template.ID})
}

func DeleteTemplate(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	var dto models.DeleteTemplateDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var template models.SquadTemplate
	err := database.DB.Where("created_by = ? AND template_name = ?", user.ID, dto.TemplateName).First(&template).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Template not found"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch template"})
		return
	}

	if err := database.DB.Where("template_id = ?", template.ID).Delete(&models.SquadUnit{}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete squad units"})
		return
	}

	if err := database.DB.Where("template_id = ?", template.ID).Delete(&models.Squad{}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete squads"})
		return
	}

	if err := database.DB.Delete(&template).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete template"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "Template deleted successfully"})
}

func FetchTemplates(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch user"})
		return
	}

	var templates []models.SquadTemplate
	if err := database.DB.Preload("Squads.Units").Where("created_by = ?", user.ID).Find(&templates).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch templates"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "templates": templates})
}
