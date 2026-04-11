package dockerpractice

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", HttpHandleAddPeople).Methods("POST")
	r.HandleFunc("/employees", HttpHandleDeletePeople).Methods("DELETE")
	r.HandleFunc("/employees", HttpHandleGetPeople).Methods("GET")

	fmt.Println("запустился сервер на порту 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}

}
