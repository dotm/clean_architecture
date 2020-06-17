package http_server

import (
	"fmt"
	"net/http"

	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/usecase/user_web"

	database "github.com/dotm/clean-architecture/backend-golang/infrastructure/database_file_storage"
	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/delivery_input/http_server_delivery"
	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/repository_output/list_repository"
	"github.com/dotm/clean-architecture/backend-golang/todo-list-microservice-v1/repository_output/user_repository"
)

//Start will start the http server after setting up the dependency tree
func Start(rootProjectDirectory string) {
	port := ":8080"
	fmt.Println("config -- http server port:", port)

	fmt.Println("initializing infrastructures...")
	database := database.NewFileStorageDatabase(rootProjectDirectory)

	fmt.Println("initializing repositories that depend on infrastructures...")
	var listRepository list_repository.Interface
	listRepository = list_repository.New(database)
	var userRepository user_repository.Interface
	userRepository = user_repository.New(database)

	fmt.Println("initializing usecases that depend on repositories through each repository interface...")
	var userWebUsecase user_web.UsecaseInterface
	userWebUsecase = user_web.New(userRepository, listRepository)

	fmt.Println("initializing deliveries that depend on usecases through each usecase interface...")
	delivery := http_server_delivery.New(userWebUsecase)
	http.HandleFunc("/ping", bypassCORS(delivery.Ping))
	http.HandleFunc("/todo-list/user-role/web-platform/v1/login", bypassCORS(delivery.UserWebLogin))
	http.HandleFunc("/todo-list/user-role/web-platform/v1/todo-lists", bypassCORS(delivery.UserWebGetAllToDoList))
	http.HandleFunc("/todo-list/user-role/web-platform/v1/todo-list/add", bypassCORS(delivery.UserWebAddNewToDoListEntry))

	fmt.Println("")
	fmt.Println("finish initializations. microservice is running on port", port)
	fmt.Println("to stop server, send interrupt signal by pressing Ctrl+C")
	var err error
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("error setting up http server", err)
	}
}

//hack for CORS using middleware because we are using localhost. DON'T USE IN PRODUCTION
func bypassCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setupHeaderForCORS(&w, r)
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	}
}
func setupHeaderForCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
