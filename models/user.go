package models 
import (
	"time"
)
type UsersPos struct {
	ID int `gorm:"primaryKey"`
	Username string  `gorm:"size:100`
	Email string  `gorm:"size:100;unique"`
	Password string 
	Code string 
	IsActive bool `gorm:"default:false"`
	CreatedAt time.Time
	Items []ItemsAdd `gorm:"foreignKey:UserID"`
}