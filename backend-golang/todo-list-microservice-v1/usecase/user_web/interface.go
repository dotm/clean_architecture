package user_web

import (
	domain "github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/domain_entities"
	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/usecase/application_entity"
)

type UsecaseInterface interface {
	Login(username string, password string) (application_entity.User, error)
	GetAllToDoList() ([]domain.List, error)
	AddNewToDoListEntry(user application_entity.User, todoList domain.List, entryText string) error
}
