package dto

type CartUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	ProductID uint64 `json:"product_id" form:"product_id" binding:"required"`
	Quantity  uint64 `json:"quantity" form:"quantity" binding:"required"`
	Price     uint64 `json:"price" form:"price" binding:"required"`
	UserID    uint64 `json:"user_id" form:"user_id" binding:"required"`
}

type CartCreateDTO struct {
	ProductID uint64 `json:"product_id" form:"product_id" binding:"required"`
	Quantity  uint64 `json:"quantity" form:"quantity" binding:"required"`
	Price     uint64 `json:"price" form:"price" binding:"required"`
	UserID    uint64 `json:"user_id" form:"user_id" binding:"required"`
}
