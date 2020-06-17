package list_repository

import domain "github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/domain_entities"

//Interface ...
type Interface interface {
	GetAll() []domain.List
	Save(list domain.List)
	//Delete(list domain.List)
}
