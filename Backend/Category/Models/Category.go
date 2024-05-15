package models

import "time"

type Category struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string    `json:"description" gorm:"type:text" validate:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}

type RequestCategoryCreate struct {
	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"required"`
}

type RequestCategoryUpdate struct {
	Name        string `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"required"`
}
