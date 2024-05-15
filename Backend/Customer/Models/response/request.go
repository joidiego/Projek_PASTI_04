package response

import "Customer/Models/entity"

type RequestCustomerRegistration struct {
	FullName    string        `json:"full_name" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Username    string        `json:"username" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Email       string        `json:"email" gorm:"type:varchar(100)" validate:"required,email"`
	Password    string        `json:"password" gorm:"type:varchar(100)" validate:"required"`
	Phone       string        `json:"phone" gorm:"type:varchar(100)" validate:"required"`
	Address     string        `json:"address" gorm:"type:varchar(100)" validate:"required"`
	Gender      entity.Gender `json:"gender" gorm:"type:varchar(20)" validate:"required"`
	DateOfBirth string        `json:"date_of_birth" gorm:"type:varchar(100)" validate:"required"`
}

type RequestCustomerLogin struct {
	Username string `json:"username" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Password string `json:"password" gorm:"type:varchar(100)" validate:"required"`
}

type RequestCustomerUpdate struct {
	FullName    string        `json:"full_name" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Username    string        `json:"username" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Email       string        `json:"email" gorm:"type:varchar(100)" validate:"required,email"`
	Phone       string        `json:"phone" gorm:"type:varchar(100)" validate:"required"`
	Address     string        `json:"address" gorm:"type:varchar(100)" validate:"required"`
	Gender      entity.Gender `json:"gender" gorm:"type:varchar(20)" validate:"required"`
	DateOfBirth string        `json:"date_of_birth" gorm:"type:varchar(100)" validate:"required"`
}

type RequestCustomerForgotPassword struct {
	Username string `json:"username" gorm:"type:varchar(100);uniqueIndex" validate:"required"`
	Phone    string `json:"phone" gorm:"type:varchar(100)" validate:"required"`
	Password string `json:"password" gorm:"type:varchar(100)" validate:"required"`
}

type RequestCustomerEditPassword struct {
	Password string `json:"password" gorm:"type:varchar(100)" validate:"required"`
}
