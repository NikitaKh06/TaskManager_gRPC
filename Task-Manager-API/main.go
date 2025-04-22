package main

import (
	task_manager_grpc_generated "github.com/NikitaKh06/TaskManagerProtoFiles/github.com/NikitaKh06/TaskManagerProtoFiles/task-manager-grpc-generated"
	"google.golang.org/grpc"
	"log"
	"task-manager-api/config"
	"task-manager-api/internal"
)

func main() {
	err := config.ConfigureLogger()
	if err != nil {
		log.Fatal(err)
	}
	
	connection, err := grpc.NewClient("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	internal.ClientgRPC = task_manager_grpc_generated.NewTaskManagerClient(connection)

}
