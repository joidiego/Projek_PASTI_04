package entity

import "time"

type Order struct {
	Id          uint      `json:"id"`
	OrderNumber string    `json:"order_number" gorm:"type:varchar(50);idxnm" validate:"required"`
	Quantity    int32     `json:"quantity" gorm:"type:int(11)" validate:"required"`
	Status      string    `json:"status" gorm:"type:string"`
	CustomerID  uint      `json:"customer_id" gorm:"idxcs"`
	ProductID   uint      `json:"product_id" gorm:"idxpd"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}
