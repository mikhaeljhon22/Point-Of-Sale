package services 

import (
	"POS/repository"
	"POS/aggregate"
)
type PersonService struct{
	person repository.Repository
}

func NewDDDService(person repository.Repository) *PersonService{
	return &PersonService{person: person,}
}

func (r *PersonService) AddService(a aggregate.Person) error{
	return r.person.Add(a)
}