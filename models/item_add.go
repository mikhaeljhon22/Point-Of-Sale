package models 
import (
	"time"
)
type ItemsAdd struct {
	ID int `gorm:"primaryKey"`
	UserID int 
	Item_name string 
	Stock int
	SKU string 
	Price string 
	Random_code string 
	CreatedAt time.Time
	User UsersPos `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
}