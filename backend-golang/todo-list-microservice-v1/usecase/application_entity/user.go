package application_entity

import (
	domain "github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/domain_entities"
)

//User ..
type User struct {
	UUID     string
	Name     string
	Password string
}

//AddEntryToList ..
func (user User) AddEntryToList(text string, list *domain.List) error {
	person := domain.Person{UUID: user.UUID, Name: user.Name}
	err := person.AddEntryToList(text, list)
	return err
}
