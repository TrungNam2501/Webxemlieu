package models

import "gorm.io/datatypes"

type Item struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
type RecipeProcessDefinition struct {
	Oid          string `gorm:"type:uuid;primaryKey"`
	RecipeID     string
	Name         string
	Type         string
	Configs      datatypes.JSON
	ProductID    string
	ProductType  string
	LimitaryHour datatypes.JSON
}

func (RecipeProcessDefinition) TableName() string {
	return "kvmes.recipe_process_definition"
}
