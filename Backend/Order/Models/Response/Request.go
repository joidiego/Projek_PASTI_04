package response

type RequestOrderCreate struct {
	Quantity int32 `json:"quantity" gorm:"type:int(11)" validate:"required"`
}
