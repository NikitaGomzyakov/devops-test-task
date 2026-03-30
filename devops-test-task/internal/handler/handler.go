package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type DataRequest struct {
	Payload string `json:"payload"`
}

type ProcessingResult struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}

type Handler struct {
	RDB *redis.Client
}

func (h *Handler) ProcessData(w http.ResponseWriter, r *http.Request) {
	var req DataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	jobID := uuid.New().String()
	ctx := context.Background()

	initial := ProcessingResult{ID: jobID, Status: "processing"}
	data, _ := json.Marshal(initial)
	h.RDB.Set(ctx, jobID, data, 1*time.Hour)

	go func(id, payload string) {
		time.Sleep(7 * time.Second)
		final := ProcessingResult{
			ID:     id,
			Status: "completed",
			Result: "Processed data: " + payload,
		}
		res, _ := json.Marshal(final)
		h.RDB.Set(context.Background(), id, res, 1*time.Hour)
	}(jobID, req.Payload)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(initial)
}

func (h *Handler) GetResults(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	val, err := h.RDB.Get(context.Background(), id).Result()
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(val))
}