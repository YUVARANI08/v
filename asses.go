package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// book stuct

type Student struct {
	ID     string  `json:"id"`
	Firstname   string  `json: "firstname"`
	Lastname  string  `json: "lastname"`
	College *College `json: "college"`
}
type College struct {
	Name string `json:"firstname"`
	Address  string `json:"lastname"`
}

var students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "appllication/json")
	json.NewEncoder(w).Encode(students)
}
func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appllication/json")

	params := mux.Vars(r)
	for _, item := range students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

		json.NewEncoder(w).Encode(&Student{})
	}

}
func updateStudentDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appllication/json")
	upid := mux.Vars(r)
	for i, val := range students {
		if val.ID == upid["id"] {
			students = append(students[:i], students[i+1:]...)
			var updatestudentdetails Student
			_ = json.NewDecoder(r.Body).Decode(&updatestudentdetails)
			updatestudentdetails.ID = upid["id"]
			students = append(students, updatestudentdetails)
			json.NewEncoder(w).Encode(updatestudentdetails)
			return
		}
	}

}
func deleteStudentDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appllication/json")

	dele := mux.Vars(r)
	for i, val := range students {
		if val.ID == dele["id"] {
			students = append(students[:i], students[i+1:]...)
			break
		}

		json.NewEncoder(w).Encode(students)

	}

}
func createStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "appllication/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(100000))
	student = append(students, student)
	json.NewEncoder(w).Encode(student)

}

func main() {
	//intial router
	router := mux.NewRouter()

	students = append(students, Student{ID: "1234", Firstname: "kavith", Lastname: "kumar", College: &College{Name: "GSS jain college", Address: "chennai,vepery,tamilnadu"}})
	students = append(students, Student{ID: "2345", Fir: "vijaya", Lastname: "shanthi", :College &College{Name: "Gss jain college", Address: "chennai,vepery,tamilnadu"}})
	router.HandleFunc("/studentss", getStudent).Methods("GET")
	router.HandleFunc("/books/{id}", getStudents).Methods("GET")
	router.HandleFunc("/books/{id}", deleteStudentDetails).Methods("DELETE")
	router.HandleFunc("/books/{id}", updateStudentDetails).Methods("PUT")
	router.HandleFunc("/book", createStudent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8083", router))
