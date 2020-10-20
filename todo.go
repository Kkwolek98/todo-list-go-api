package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/gorilla/mux"
)

var db *gorm.DB
var err error
var dsn = "user=postgres password=admin dbname=todo port=5432 sslmode=disable"

// Todo type
type Todo struct {
	gorm.Model
	Title 			string 	`json:"title"`
	Description 	string 	`json:"description"`
	IsDone 			bool 	`json:"isDone"`
}

func initialMigration() {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connecto to DB")
	}

	db.AutoMigrate(&Todo{})
}

//Functions
func getTodos(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connecto to DB")
	}

	var todos []Todo
	db.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connecto to DB")
	}

	id := mux.Vars(r)["id"]

	var todo Todo
	db.Where("id = ?", id).Find(&todo)

	json.NewEncoder(w).Encode(&todo)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connecto to DB")
	}
	
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	
	db.Create(&todo)

	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connecto to DB")
	}
	
	id := mux.Vars(r)["id"]

	var todo Todo
	db.Where("id = ?", id).Find(&todo)
	err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	db.Save(&todo)

	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connecto to DB")
	}
	
	id := mux.Vars(r)["id"]

	var todo Todo
	db.Where("id = ?", id).Find(&todo)
	db.Delete(&todo)

	json.NewEncoder(w).Encode(todo)
}

func toggleTodoDone(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connecto to DB")
	}
	
	id := mux.Vars(r)["id"]

	var todo Todo
	db.Where("id = ?", id).Find(&todo)

	todo.IsDone = !todo.IsDone

	db.Save(&todo)

	json.NewEncoder(w).Encode(todo)
}