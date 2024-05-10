package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	empHTTP "algogrit.com/emp_server/employees/http"
	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/employees/service"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	var repo = repository.NewInMem()
	var v1Svc = service.NewV1(repo)
	var empHandler = empHTTP.New(v1Svc)

	empHandler.SetupRoutes(r)

	log.Println("Starting server on port: 8000...")

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
