package entity

import (
	"gorm.io/gorm"
)

type Slider struct {
	gorm.Model
	ID    uint64 `gorm:"primary_key:auto_increment"` // set primary key dan auto increment untuk kolom id
	Name  string `gorm:"type:varchar(255)" json:"name" validate:"required,min=3,max=255"`
	Image string `gorm:"type:varchar(255)" json:"image" validate:"required,oneof=jpg jpeg png"`
}
