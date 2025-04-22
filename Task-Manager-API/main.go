package main

import (
	task_manager_grpc_generated "github.com/NikitaKh06/TaskManagerProtoFiles/github.com/NikitaKh06/TaskManagerProtoFiles/task-manager-grpc-generated"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"task-manager-api/config"
	"task-manager-api/internal"
)

func main() {
	log.SetOutput(os.Stderr)
	err := config.ConfigureLogger()
	if err != nil {
		log.Fatal(err)
	}

	connection, err := grpc.Dial("task-manager-database-service-grpc:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	internal.ClientgRPC = task_manager_grpc_generated.NewTaskManagerClient(connection)

	http.HandleFunc("/create", internal.CreateHandler)
	http.HandleFunc("/list", internal.ListHandler)
	http.HandleFunc("/delete", internal.DeleteHandler)
	http.HandleFunc("/done", internal.DoneHandler)

	http.ListenAndServe(":8081", nil)
}
