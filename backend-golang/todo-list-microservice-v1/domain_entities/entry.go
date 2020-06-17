package domain

//domain should only depends on other domain entities

//Entry represents an entry in a to-do list.
type Entry struct {
	UUID      string
	Text      string
	Owner Person
}

const maximumEntryLength int = 50
