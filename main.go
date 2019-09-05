package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Worlds")
}

func handleRequests() {
	//using gorilla mux
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", helloWorld).Methods("GET")

	//users endpoint
	r.HandleFunc("/users", AllUsers).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	fmt.Println("Go run on port 8081")

	//db operation
	Migrate()

	//run the func
	handleRequests()
}
