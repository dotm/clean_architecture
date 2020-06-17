package domain

//domain should only depends on other domain entities

//List represents a real world to do list
type List struct {
	UUID    string
	Title   string
	Entries []Entry
}
