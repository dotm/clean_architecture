package user_web

import (
	"fmt"

	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/repository_output/list_repository"

	domain "github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/domain_entities"
	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/repository_output/user_repository"
	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/usecase/application_entity"
)

//UsecaseImplementation ..
type UsecaseImplementation struct {
	userRepository user_repository.Interface
	listRepository list_repository.Interface
}

//New ..
func New(userRepository user_repository.Interface, listRepository list_repository.Interface) UsecaseImplementation {
	return UsecaseImplementation{
		userRepository: userRepository,
		listRepository: listRepository,
	}
}

//Login ..
func (usecase UsecaseImplementation) Login(username string, password string) (application_entity.User, error) {
	emptyUser := application_entity.User{}

	result := usecase.userRepository.CheckUserExistAndPasswordCorrect(username, password)
	if result == false {
		return emptyUser, fmt.Errorf("wrong username or password")
	}

	user := usecase.userRepository.GetUserByUserName(username)
	return user, nil
}

//GetAllToDoList ..
func (usecase UsecaseImplementation) GetAllToDoList() ([]domain.List, error) {
	allToDoList := usecase.listRepository.GetAll()
	return allToDoList, nil
}

//AddNewToDoListEntry ..
func (usecase UsecaseImplementation) AddNewToDoListEntry(user application_entity.User, todoList domain.List, entryText string) error {
	err := user.AddEntryToList(entryText, &todoList)
	if err != nil {
		return err
	}

	usecase.listRepository.Save(todoList)
	return nil
}
