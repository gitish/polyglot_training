package models

import "gorm.io/gorm"

type User struct {
	gorm.Model         // GORM provides a predefined struct named gorm.Model, which includes commonly used fields like ID, CreatedAt, UpdatedAt etc.
	Name       string  // A regular string field
	Email      *string // A pointer to a string, allowing for null values
}
