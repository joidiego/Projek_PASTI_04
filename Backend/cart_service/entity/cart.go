package entity

type Cart struct {
	ID        uint64 `gorm:"primary_key:auto_increment"`
	ProductID uint64 `gorm:"type:int(11)"`
	Quantity  uint64 `gorm:"type:int(11)"`
	Price     uint64 `gorm:"type:int(11)"`
	UserID    uint64 `gorm:"type:int(11)"`
}
