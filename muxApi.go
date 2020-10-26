package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var students []Student

func getstudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //fetching params
	for _, item := range students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Student Student
	_ = json.NewDecoder(r.Body).Decode(&Student)
	count := len(students)
	Student.ID = strconv.Itoa(count + 1)
	students = append(students, Student)
	json.NewEncoder(w).Encode(Student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var Student Student
			_ = json.NewDecoder(r.Body).Decode(&Student)
			Student.ID = params["id"]
			students = append(students, Student)
			json.NewEncoder(w).Encode(Student)
			return
		}
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}
func main() {

	r := mux.NewRouter()

	students = append(students, Student{ID: "1", Name: "Student One"})
	students = append(students, Student{ID: "2", Name: "Student Two"})

	r.HandleFunc("/students", getstudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
