package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// type Address struct {
// 	City string `json:"city"`
// }

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"-"`
	// Address Address `json:"address"`
}

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"id": %d, "name": "%s", "speciality": "%s"}`, e.ID, e.Name, e.Department)

// 	return []byte(jsonString), nil
// }

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Shikhar", "Cloud", 10002},
	{3, "Mark", "SRE", 2003},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmp.ID = len(employees) + 1
	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
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

	// http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
	http.ListenAndServe("localhost:8000", handlers.LoggingHandler(os.Stdout, r))
}
