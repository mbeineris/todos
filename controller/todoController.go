package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"todos/errorhandler"
	"todos/model"
)

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	// Create instance of Todo based on form input
	todo := model.Todo{
		Name: r.FormValue("name"),
	}

	// Convert id<string> from request to id<int>
	tododata_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(tododata_id)
	}

	// Handle DB changes
	dbErr, newTodo := model.TodoUpdate("TodoBucket", tododata_id, todo)
	if dbErr == nil {

		todo, err := json.Marshal(newTodo)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(todo)
	}
}

func TodoRemove(w http.ResponseWriter, r *http.Request) {

	// Convert id<string> from request to id<int>
	todolist_id, err := strconv.Atoi(mux.Vars(r)["todolist_id"])
	if err != nil {
		log.Fatal(todolist_id)
	}

	todo_id, err := strconv.Atoi(mux.Vars(r)["todo_id"])
	if err != nil {
		log.Fatal(todo_id)
	}

	// Handle DB changes
	dbErr := model.TodoRemove("TodoBucket", todolist_id, todo_id)
	if dbErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("200 - Remove Successful!"))
	}

}

func TodoToggle(w http.ResponseWriter, r *http.Request) {

	// Convert id<string> from request to id<int>
	todolist_id, err := strconv.Atoi(mux.Vars(r)["todolist_id"])
	if err != nil {
		log.Fatal(todolist_id)
	}

	todo_id, err := strconv.Atoi(mux.Vars(r)["todo_id"])
	if err != nil {
		log.Fatal(todo_id)
	}

	status, err := strconv.ParseBool(r.FormValue("status"))
	if err != nil {
		log.Fatal(status)
	}

	// Handle DB changes
	dbErr := model.TodoToggle("TodoBucket", todolist_id, todo_id, status)
	if dbErr != nil {
		errorhandler.CatchError(dbErr, "Todo not found.")
	}

}
