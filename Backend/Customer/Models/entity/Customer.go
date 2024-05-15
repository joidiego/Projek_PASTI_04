package entity

import "time"

type Customer struct {
	Id          uint      `json:"id"`
	FullName    string    `json:"full_name" gorm:"type:varchar(100)" validate:"required"`
	Username    string    `json:"username" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Email       string    `json:"email" gorm:"type:varchar(100);unique" validate:"required,email"`
	Password    string    `json:"password" gorm:"type:varchar(100)" validate:"required"`
	Phone       string    `json:"phone" gorm:"type:varchar(100)" validate:"required"`
	Address     string    `json:"address" gorm:"type:varchar(100)" validate:"required"`
	Gender      Gender    `json:"gender" gorm:"type:varchar(20)" validate:"required"`
	DateOfBirth string    `json:"date_of_birth" gorm:"type:varchar(100)" validate:"required"`
	Image       string    `json:"image" gorm:"type:varchar(100)" validate:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}

type Gender string

const (
	LakiLaki  Gender = "Laki-Laki"
	Perempuan Gender = "Perempuan"
)
