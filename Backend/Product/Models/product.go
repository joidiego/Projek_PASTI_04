package models

import (
	"time"
)

type Product struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(50);index:idx_nm,unique" validate:"required"`
	Description string    `json:"description" gorm:"type:text" validate:"required"`
	Image       string    `json:"image" gorm:"type:varchar(50)" validate:"required"`
	Price       string    `json:"price" gorm:"type:varchar(50)" validate:"required"`
	CategoryID  uint      `json:"category_id" gorm:"index"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}

type RequestProductCreate struct {
	Name        string `json:"name" gorm:"type:varchar(50);index:idx_nm,unique" validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"required"`
	Price       string `json:"price" gorm:"type:varchar(50)" validate:"required"`
}

type RequestProductUpdate struct {
	Name        string `json:"name" gorm:"type:varchar(50);index:idx_nm,unique" validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"required"`
	Price       string `json:"price" gorm:"type:varchar(50)" validate:"required"`
	CategoryID  uint   `json:"category_id" gorm:"index"`
}

type Category struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Description string    `json:"description" gorm:"type:text" validate:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
