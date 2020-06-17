package list_repository

import (
	"encoding/json"
	"fmt"

	//aside from domain, repository is allowed to depend on infrastructure
	database "github.com/dotm/clean-architecture/backend-golang/infrastructure/database_file_storage"
	domain "github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/domain_entities"
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

const toDoListPrefix = "todo_list"

func (repository Struct) GetAll() []domain.List {
	allTodoListFilePath, err := repository.Database.GetAllFilePathWithPrefix(toDoListPrefix)
	if err != nil {
		fmt.Println("failed getting all todo list file path:", err)
	}

	allToDoList := []domain.List{}
	for _, filePath := range allTodoListFilePath {
		bytes := repository.Database.ReadFromFile(filePath)
		var todoList domain.List
		err = json.Unmarshal(bytes, &todoList)
		if err != nil {
			fmt.Println("failed to get unmarshall JSON from file:", filePath)
			continue
		}
		allToDoList = append(allToDoList, todoList)
	}

	return allToDoList
}

//Save will save the list to database
func (repository Struct) Save(list domain.List) {
	bytes, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON", err)
	}
	filename := toDoListPrefix + "---title__" + list.Title + "---uuid__" + list.UUID + ".json"
	repository.Database.CreateNewOrOverwriteExistingFileWithBytes(filename, bytes)
}
