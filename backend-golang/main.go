package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/dotm/clean-architecture/backend-golang/microservice-server/http_server"
)

func main() {
	//decide which server to run based on the first argument passed to command line
	firstArgumentUnspecified := len(os.Args) < 2
	var firstArgumentOfCommand string
	if firstArgumentUnspecified {
		firstArgumentOfCommand = "http_server"
	} else {
		firstArgumentOfCommand = os.Args[1]
	}

	fmt.Println("initializing microservice configuration...")
	rootProjectDirectory := getRootProjectDirectory()
	fmt.Println("config -- root project directory:", rootProjectDirectory)

	switch firstArgumentOfCommand {
	case "http_server":
		http_server.Start(rootProjectDirectory)
	default:
		panic("server unknown: " + firstArgumentOfCommand)
	}
}

func getRootProjectDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	return basepath
}
