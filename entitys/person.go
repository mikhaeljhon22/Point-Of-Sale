package entitys
import (
	"time"
)

type Person struct {
	ID int `gorm:"primaryKey"`
	Name string 
	Age int
	CreatedAt time.Time
}