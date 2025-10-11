package memory 
import (
	"gorm.io/gorm"
	"POS/aggregate"
	"POS/entitys"
)

type Memory struct {
	db *gorm.DB
}

func NewMemory(db *gorm.DB) *Memory{
	return &Memory{db:db,}
}
func (m *Memory) Add(c aggregate.Person)error{
	name := c.GetPerson()
	age := c.GetAge()
	data := entitys.Person{
		Name: name,
		Age: age,
	}
	return m.db.Create(&data).Error
}