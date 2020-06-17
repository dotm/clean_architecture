package http_server_delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	domain "github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/domain_entities"
	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/usecase/application_entity"
)

//UserWebLogin ..
func (delivery Delivery) UserWebLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receive request on UserWebLogin")

	//parse data out of request
	var jsonMap map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&jsonMap)
	username, _ := jsonMap["username"].(string)
	password, _ := jsonMap["password"].(string)

	//call usecase
	user, err := delivery.userWebUsecase.Login(username, password)

	//write response
	data := map[string]interface{}{
		"user": user,
	}
	writeResponse(w, data, err)
}

//UserWebGetAllToDoList ..
func (delivery Delivery) UserWebGetAllToDoList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receive request on UserWebGetAllToDoList")

	//parse data out of request
	//nothing to parse...

	//call usecase
	allToDoList, err := delivery.userWebUsecase.GetAllToDoList()

	//write response
	data := map[string]interface{}{
		"allToDoList": allToDoList,
	}
	writeResponse(w, data, err)
}

//UserWebAddNewToDoListEntry ..
func (delivery Delivery) UserWebAddNewToDoListEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receive request on UserWebAddNewToDoListEntry")

	//parse data out of request
	var jsonObj todoListAddEntryRequest
	_ = json.NewDecoder(r.Body).Decode(&jsonObj)

	//call usecase
	err := delivery.userWebUsecase.AddNewToDoListEntry(jsonObj.User, jsonObj.TodoList, jsonObj.EntryText)

	//write response
	data := map[string]interface{}{}
	writeResponse(w, data, err)
}

type todoListAddEntryRequest struct {
	User      application_entity.User
	TodoList  domain.List
	EntryText string
}
