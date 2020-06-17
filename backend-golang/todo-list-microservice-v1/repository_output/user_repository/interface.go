package user_repository

import "github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/usecase/application_entity"

//Interface ..
type Interface interface {
	CheckUserExistAndPasswordCorrect(username string, password string) bool
	GetUserByUserName(username string) application_entity.User
}
