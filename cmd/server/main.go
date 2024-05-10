package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	empRepo "algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/entities"
)

var repo = empRepo.NewInMem()

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	employees, err := repo.ListAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := repo.Save(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else if req.Method == "GET" {
// 		EmployeesIndexHandler(w, req)
// 	} else {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	h := func(w http.ResponseWriter, req *http.Request) {
// 		begin := time.Now()

// 		next.ServeHTTP(w, req)

// 		fmt.Printf("%s %s took %s\n", req.Method, req.URL, time.Since(begin))
// 	}

// 	return http.HandlerFunc(h)
// }

func main() {
	r := mux.NewRouter()
	// r := http.NewServeMux()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	// r.HandleFunc("/employees", EmployeesHandler)
	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	log.Println("Starting server on port: 8000...")
	// http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
