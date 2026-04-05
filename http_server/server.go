package httpserver

import (
	"errors"
	"fmt"
	"net/http"
)

func StartHttpServer() error {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("патерн", r.URL.Path)
		w.Write([]byte("Hello from docker\n"))
	})

	fmt.Println("Server is listening...")

	err := http.ListenAndServe(":8000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
