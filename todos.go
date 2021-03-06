package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"todos/controller"
	"todos/database"
)

// Core function
func main() {

	database.Connect()

	// Initialize Router
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/get_todolist", controller.GetTodoList).Methods("GET")
	api.HandleFunc("/create", controller.TodoListCreate).Methods("POST")
	api.HandleFunc("/remove", controller.RemoveTodoList).Methods("POST")
	api.HandleFunc("/{id}/create", controller.TodoCreate).Methods("POST")
	api.HandleFunc("/{todolist_id}/todo/{todo_id}/remove", controller.TodoRemove).Methods("POST")
	api.HandleFunc("/{todolist_id}/todo/{todo_id}/toggle", controller.TodoToggle).Methods("PATCH")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("view/dist/")))
	r.PathPrefix("/").HandlerFunc(controller.GetTodoList)

	http.ListenAndServe(":3000", handlers.CompressHandler(r))

	// For self-signed certificate
	// err := http.ListenAndServeTLS(":3000", "ssl/server.crt", "ssl/server.key", handlers.CompressHandler(r))
	// if err != nil {
	// 	panic(err)
	// }
}
