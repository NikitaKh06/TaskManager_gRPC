package internal

import (
	"context"
	task_manager_grpc_generated "github.com/NikitaKh06/TaskManagerProtoFiles/task-manager-grpc-generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"task-manager-database/config"
)

type TaskManagerService interface {
	CreateTask(ctx context.Context, request *task_manager_grpc_generated.CreateTaskRequest) error
	GetTask(ctx context.Context) (*task_manager_grpc_generated.GetTaskResponse, error)
	DeleteTask(ctx context.Context) error
	DoneTask(ctx context.Context) error
}

type taskManagerServerAPI struct {
	task_manager_grpc_generated.UnimplementedTaskManagerServer
	taskManagerService TaskManagerService
}

func Register(gRPCServer *grpc.Server, service TaskManagerService) {
	task_manager_grpc_generated.RegisterTaskManagerServer(gRPCServer, &taskManagerServerAPI{taskManagerService: service})
}

func (*taskManagerServerAPI) CreateTask(ctx context.Context, request *task_manager_grpc_generated.CreateTaskRequest) error {
	log.Println("Request for creating new task")

	if request.Name == "" {
		log.Println("Error: name is empty")
		return status.Error(codes.InvalidArgument, "Name must be not null")
	}

	query := "INSERT INTO tasks (name, text) VALUES ($1, $2)"

	_, err := config.Db.Exec(context.Background(), query, request.Name, request.Text)
	if err != nil {
		log.Println("Error: insertion into database")
		return status.Error(codes.Internal, "Can not insert task into database")
	}

	log.Println("Request completed")
	return nil
}

func (*taskManagerServerAPI) GetTask(ctx context.Context, request *task_manager_grpc_generated.GetTaskRequest) (*task_manager_grpc_generated.GetTaskResponse, error) {
	log.Println("Request for list of tasks")

	//query := "SELECT * FROM tasks"

	return nil, nil
}

func (*taskManagerServerAPI) DeleteTask(ctx context.Context, request *task_manager_grpc_generated.DeleteTaskRequest) error {
	log.Println("Request for deleting the task")

	query := "DELETE FROM tasks WHERE ID = $1"

	_, err := config.Db.Exec(context.Background(), query, request.Id)
	if err != nil {
		log.Println("Error: deletion from database")
		return status.Error(codes.Internal, "Can not delete task from database")
	}

	log.Println("Request completed")
	return nil
}

func (*taskManagerServerAPI) DoneTask(ctx context.Context, request *task_manager_grpc_generated.DoneTaskRequest) error {
	log.Println("Request for marking the task as Done")

	query := "UPDATE tasks SET done = TRUE WHERE id = $1"

	_, err := config.Db.Exec(context.Background(), query, request.Id)
	if err != nil {
		log.Println("Error: marking the task as Done")
		return status.Error(codes.Internal, "Can not mark task as Done")
	}

	log.Println("Request completed")
	return nil
}
