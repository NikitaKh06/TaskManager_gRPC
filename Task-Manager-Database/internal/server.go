package internal

import (
	"context"
	task_manager_grpc_generated "github.com/NikitaKh06/TaskManagerProtoFiles/github.com/NikitaKh06/TaskManagerProtoFiles/task-manager-grpc-generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"task-manager-database/config"
)

var gRPCServer *grpc.Server

func RunApp() error {
	log.Println("Request for running the app")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}

	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Println("Error" + err.Error())
		return err
	}

	return nil
}

func CreateServer() {
	log.Println("Request for creating new server")
	gRPCServer = grpc.NewServer()
	Register(gRPCServer)
	log.Println("Request completed")
}

type TaskManagerService interface {
	CreateTask(ctx context.Context, request *task_manager_grpc_generated.CreateTaskRequest) (*emptypb.Empty, error)
	GetTask(ctx context.Context, request *task_manager_grpc_generated.GetTaskRequest) (*task_manager_grpc_generated.GetTaskResponse, error)
	DeleteTask(ctx context.Context, request *task_manager_grpc_generated.DeleteTaskRequest) (*emptypb.Empty, error)
	DoneTask(ctx context.Context, request *task_manager_grpc_generated.DoneTaskRequest) (*emptypb.Empty, error)
}

type taskManagerServerAPI struct {
	task_manager_grpc_generated.UnimplementedTaskManagerServer
}

func Register(gRPCServer *grpc.Server) {
	task_manager_grpc_generated.RegisterTaskManagerServer(gRPCServer, &taskManagerServerAPI{})
}

func (s *taskManagerServerAPI) CreateTask(ctx context.Context, request *task_manager_grpc_generated.CreateTaskRequest) (*emptypb.Empty, error) {
	log.Println("Request for creating new task")

	if request.Name == "" {
		log.Println("Error: name is empty")
		return nil, status.Error(codes.InvalidArgument, "Name must be not null")
	}

	query := "INSERT INTO tasks (name, text) VALUES ($1, $2)"

	_, err := config.Db.Exec(context.Background(), query, request.Name, request.Text)
	if err != nil {
		log.Println("Error: insertion into database")
		return nil, status.Error(codes.Internal, "Can not insert task into database")
	}

	log.Println("Request completed")
	return &emptypb.Empty{}, nil
}

func (s *taskManagerServerAPI) GetTask(ctx context.Context, request *task_manager_grpc_generated.GetTaskRequest) (*task_manager_grpc_generated.GetTaskResponse, error) {
	log.Println("Request for list of tasks")

	query := "SELECT * FROM tasks"
	rows, err := config.Db.Query(context.Background(), query)
	if err != nil {
		log.Println("Error: " + err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	var allTasks []*task_manager_grpc_generated.Task

	for rows.Next() {
		var task task_manager_grpc_generated.Task
		err = rows.Scan(&task.Id, &task.Name, &task.Text, &task.Done)
		if err != nil {
			log.Println("Error: " + err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}

		allTasks = append(allTasks, &task)
	}

	response := task_manager_grpc_generated.GetTaskResponse{
		Tasks: allTasks,
	}

	log.Println("Request completed")
	return &response, nil
}

func (s *taskManagerServerAPI) DeleteTask(ctx context.Context, request *task_manager_grpc_generated.DeleteTaskRequest) (*emptypb.Empty, error) {
	log.Println("Request for deleting the task")

	query := "DELETE FROM tasks WHERE id = $1"

	action, err := config.Db.Exec(context.Background(), query, request.Id)
	if err != nil {
		log.Println("Error: deletion from database")
		return nil, status.Error(codes.Internal, "Can not delete task from database")
	}

	if action.RowsAffected() == 0 {
		log.Println("Request canceled: Task not found")
		return nil, status.Error(codes.NotFound, "Task not found")
	}

	log.Println("Request completed")
	return &emptypb.Empty{}, nil
}

func (s *taskManagerServerAPI) DoneTask(ctx context.Context, request *task_manager_grpc_generated.DoneTaskRequest) (*emptypb.Empty, error) {
	log.Println("Request for marking the task as Done")

	query := "UPDATE tasks SET done = TRUE WHERE id = $1"

	action, err := config.Db.Exec(context.Background(), query, request.Id)
	if err != nil {
		log.Println("Error: marking the task as Done")
		return nil, status.Error(codes.Internal, "Can not mark task as Done")
	}

	if action.RowsAffected() == 0 {
		log.Println("Request canceled: Task not found")
		return nil, status.Error(codes.NotFound, "Task not found")
	}

	log.Println("Request completed")
	return &emptypb.Empty{}, nil
}
