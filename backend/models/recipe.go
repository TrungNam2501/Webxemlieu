package models

import "gorm.io/datatypes"

// RecipeProcessDefinition ánh xạ tới bảng kvmes.recipe_process_definition trong KVERP.
type RecipeProcessDefinition struct {
	Oid          string         `json:"Oid" gorm:"type:uuid;primaryKey"`
	RecipeID     string         `json:"RecipeID"`
	Name         string         `json:"Name"`
	Type         string         `json:"Type"`
	Configs      datatypes.JSON `json:"Configs"`
	ProductID    string         `json:"ProductID"`
	ProductType  string         `json:"ProductType"`
	LimitaryHour datatypes.JSON `json:"LimitaryHour"`
}

// TableName trả về tên bảng đầy đủ (gồm schema) trong PostgreSQL.
func (RecipeProcessDefinition) TableName() string {
	return "kvmes.recipe_process_definition"
}
