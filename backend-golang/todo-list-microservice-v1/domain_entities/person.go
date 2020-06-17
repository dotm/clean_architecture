package domain

import (
	"fmt"

	"github.com/dotm/clean-architecture/backend-golang/utilities/uuid"
)

//domain should only depends on other domain entities

//Person is a representation of a real world human.
//This is equivalent to buyer, seller, player, customer, etc.
type Person struct {
	UUID string
	Name string
}

//CreateNewList will create a new to-do list
func (person *Person) CreateNewList(title string) List {
	return List{
		UUID:    uuid.New(),
		Title:   title,
		Entries: []Entry{},
	}
}

//AddEntryToList will add an entry to a to-do list
//if it fulfills certain restrictions (business logic)
func (person Person) AddEntryToList(text string, list *List) error {
	//the business logic
	if len(text) > maximumEntryLength {
		return fmt.Errorf("maximum allowed entry text exceeded")
	}

	entry := Entry{
		UUID:  uuid.New(),
		Text:  text,
		Owner: person,
	}
	list.Entries = append(list.Entries, entry)
	return nil
}
