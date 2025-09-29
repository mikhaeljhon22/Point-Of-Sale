package models 
import (
	"time"
)
type UsersPos struct {
	ID int `gorm:"primaryKey"`
	Username string  `gorm:"size:100`
	Email string  `gorm:"size:100;unique"`
	CreatedAt time.Time
}