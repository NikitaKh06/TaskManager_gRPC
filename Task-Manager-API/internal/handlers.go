package internal

import (
	"encoding/json"
	task_manager_grpc_generated "github.com/NikitaKh06/TaskManagerProtoFiles/github.com/NikitaKh06/TaskManagerProtoFiles/task-manager-grpc-generated"
	"log"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for creating new task")
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var task CreateTask
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var gRPCTask task_manager_grpc_generated.CreateTaskRequest
	gRPCTask.Name = task.Name
	gRPCTask.Text = task.Text

	_, err = ClientgRPC.CreateTask(r.Context(), &gRPCTask)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Request completed")
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for getting all tasks")
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	response, err := ClientgRPC.GetTask(r.Context(), &task_manager_grpc_generated.GetTaskRequest{})
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(response.Tasks)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	log.Println("Request completed")
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for deleting the task")
	if r.Method != http.MethodDelete {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var task TaskById
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var gRPCTask task_manager_grpc_generated.DeleteTaskRequest
	gRPCTask.Id = int64(task.Id)

	_, err = ClientgRPC.DeleteTask(r.Context(), &gRPCTask)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Request completed")
}

func DoneHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for marking the task as done")
	if r.Method != http.MethodPut {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var task TaskById
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var gRPCTask task_manager_grpc_generated.DoneTaskRequest
	gRPCTask.Id = int64(task.Id)

	_, err = ClientgRPC.DoneTask(r.Context(), &gRPCTask)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Request completed")
}
