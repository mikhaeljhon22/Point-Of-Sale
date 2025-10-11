package repository
import (
	"POS/aggregate"
)
type Repository interface{
	Add(aggregate.Person) error
}