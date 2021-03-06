package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//Middleware
func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}


func main() {
	initialMigration()

	// router init
	router := mux.NewRouter()

	// middleware
	router.Use(contentTypeMiddleware)

	// route handlers
	router.HandleFunc("/api/todos", getTodos).Methods("GET")
	router.HandleFunc("/api/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/api/todo", createTodo).Methods("POST")
	router.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/api/todo/toggle_done/{id}", toggleTodoDone).Methods("PUT")
	router.HandleFunc("/api/todo/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}