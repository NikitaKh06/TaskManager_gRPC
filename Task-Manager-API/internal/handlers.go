package internal

import (
	"encoding/json"
	"fmt"
	task_manager_grpc_generated "github.com/NikitaKh06/TaskManagerProtoFiles/github.com/NikitaKh06/TaskManagerProtoFiles/task-manager-grpc-generated"
	"log"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	response, err := ClientgRPC.GetTask(r.Context(), &task_manager_grpc_generated.GetTaskRequest{})
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(response.Tasks)

	jsonResponse, err := json.Marshal(response.Tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

}

func DoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

}
