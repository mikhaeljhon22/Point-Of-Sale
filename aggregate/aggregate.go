package aggregate
import (
	"POS/entitys"
	"POS/valueObject"
)
type Person struct {
	persons  *entitys.Person
	addresss valueObject.Address
}

func NewAggregate(name string, age int, addressName string, city string) (Person){
	person := &entitys.Person{
		Name: name,
		Age: age,
	}
	address := valueObject.Address{
		Address: addressName,
		City: city,
	}

	return Person{
		persons: person,
		addresss: address,
	}
}

func (p Person) GetPerson() string {
	return p.persons.Name
}
func (p Person) GetAge() int {
	return p.persons.Age
}

func (p Person) GetAddress() string {
	return p.addresss.Address
}

func (p Person) GetCity() string {
	return p.addresss.City
}