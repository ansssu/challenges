package main

import (
	"log"
	"net/http"
)

type userHandler struct{}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/users", &userHandler{})
	log.Println("HTTP Server is ready and Listening...")
	http.ListenAndServe(":3000", mux)
}
