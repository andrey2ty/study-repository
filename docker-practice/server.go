package dockerpractice

import (
	"context"
	"fmt"
	"net/http"
	"study/postgres"

	"github.com/gorilla/mux"
)

func Server() {
	ctx := context.Background()
	conn, err := postgres.ConnetctionToPostgres(ctx)

	if err != nil {
		fmt.Println("ошибка подключения к бд")
		return
	}

	h := NewHttpHandler(ctx, conn)

	r := mux.NewRouter()

	r.HandleFunc("/employees", h.HttpHandleAddPeople).Methods("POST")
	r.HandleFunc("/employees", h.HttpHandleDeletePeople).Methods("DELETE")
	r.HandleFunc("/employees", h.HttpHandleGetPeople).Methods("GET")

	fmt.Println("запустился сервер на порту 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}
