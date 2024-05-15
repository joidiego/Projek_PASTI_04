package entity

import "time"

type Vendor struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Phone     string    `json:"phone" gorm:"type:varchar(255)" validate:"required"`
	Username  string    `json:"username" gorm:"type:varchar(255)" validate:"required,min=5"`
	Password  string    `json:"password" gorm:"type:varchar(255)" validate:"required,min=5"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
