package user_repository

import (
	"encoding/json"
	"fmt"

	//aside from domain, repository is allowed to depend on infrastructure
	database "github.com/dotm/clean-architecture/backend-golang/infrastructure/database_file_storage"
	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/usecase/application_entity"
)

//Struct ..
type Struct struct {
	Database database.FileStorageDatabase
}

//New will create a new list repository
func New(database database.FileStorageDatabase) Struct {
	return Struct{
		Database: database,
	}
}

const userFilename string = "users.json"

func (repository Struct) getExistingUsers() []application_entity.User {
	var userList []application_entity.User
	//get existing users
	bytes := repository.Database.ReadFromFile(userFilename)
	err := json.Unmarshal(bytes, &userList)
	if err != nil {
		fmt.Println("error decoding JSON", err)
	}
	return userList
}

//CheckUserExistAndPasswordCorrect ..
func (repository Struct) CheckUserExistAndPasswordCorrect(username string, password string) bool {
	userList := repository.getExistingUsers()
	for _, user := range userList {
		if user.Name == username && user.Password == password {
			return true
		}
	}
	return false
}

//GetUserByUserName ..
func (repository Struct) GetUserByUserName(username string) application_entity.User {
	userList := repository.getExistingUsers()
	for _, user := range userList {
		if user.Name == username {
			return user
		}
	}
	return application_entity.User{}
}
