package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NikitaGomzyakov/devops-test/internal/handler"
	"github.com/NikitaGomzyakov/devops-test/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := storage.NewRedisClient(redisAddr)
	h := &handler.Handler{RDB: rdb}

	r := mux.NewRouter()
	r.HandleFunc("/process-data", h.ProcessData).Methods("POST")
	r.HandleFunc("/results/{id}", h.GetResults).Methods("GET")

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
