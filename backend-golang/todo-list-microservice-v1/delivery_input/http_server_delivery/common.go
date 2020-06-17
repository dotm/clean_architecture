package http_server_delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/usecase/user_web"
)

//Delivery ..
type Delivery struct {
	userWebUsecase user_web.UsecaseInterface
}

//New ..
func New(userWebUsecase user_web.UsecaseInterface) Delivery {
	return Delivery{
		userWebUsecase: userWebUsecase,
	}
}

//Ping the server
func (delivery Delivery) Ping(w http.ResponseWriter, r *http.Request) {
	pong := "pong"
	fmt.Println(pong)

	data := map[string]string{
		"data": pong,
	}
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("failed marshalling json from ping")
	}

	_, err = w.Write(bytes)
	if err != nil {
		fmt.Println("failed writing response from ping")
	}
}

func writeResponse(w http.ResponseWriter, data interface{}, err error) {
	var response map[string]interface{}
	if err != nil {
		response = map[string]interface{}{
			"error": err.Error(),
			"data":  nil,
		}
	} else {
		response = map[string]interface{}{
			"error": nil,
			"data":  data,
		}
	}

	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("failed marshalling json")
	}

	_, err = w.Write(bytes)
	if err != nil {
		fmt.Println("failed writing response")
	}
}
