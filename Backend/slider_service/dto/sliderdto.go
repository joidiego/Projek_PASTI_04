package dto

type SliderUpdateDTO struct {
	ID    uint64 `json:"id" form:"id" binding:"required"`
	Name  string `gorm:"type:varchar(255)" json:"name" validate:"required,min=3,max=255"`
	Image string `gorm:"type:varchar(255)" json:"image" validate:"required,oneof=jpg jpeg png"`
}

type SliderCreateDTO struct {
	Name  string `gorm:"type:varchar(255)" json:"name" validate:"required,min=3,max=255"`
	Image string `gorm:"type:varchar(255)" json:"image" validate:"required,oneof=jpg jpeg png"`
}
