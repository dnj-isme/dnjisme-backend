package models

import "gorm.io/gorm"

// ####################### General #######################

type Role struct {
	gorm.Model
	RoleName string `gorm:"uniqueIndex"`
}

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	DiscordID string
	RoleID    uint

	Role Role `gorm:"foreignKey:RoleID"`
}

type Portfolio struct {
	gorm.Model
	RepositoryLink string
	DemoLink       string
	Title          string
	Details        string
	MediaLink      string
	Tags           string
}

// ####################### Grandivory Template #######################

type SquadTemplate struct {
	gorm.Model
	CreatedBy    uint   `gorm:"uniqueIndex:idx_template_name_created_by"`
	TemplateName string `gorm:"uniqueIndex:idx_template_name_created_by"`
	TemplateType string

	CreatedByUser User    `gorm:"foreignKey:CreatedBy"`
	Squads        []Squad `gorm:"foreignKey:TemplateID"`
}

type Squad struct {
	gorm.Model
	SquadName      string
	TemplateID     uint
	SquadWeightGAC float64
	SquadWeightTW  float64

	Template SquadTemplate `gorm:"foreignKey:TemplateID"`
	Units    []SquadUnit   `gorm:"foreignKey:SquadID"`
}

type SquadUnit struct {
	gorm.Model
	SquadID        uint
	CharacterID    string
	UnitWeight     float64
	UnitOrder      int
	PriorityNumber int

	Squad Squad `gorm:"foreignKey:SquadID"`
}
