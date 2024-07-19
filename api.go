package main

import (
	"context"
	"encoding/json"
	"net/http"
	// "github.com/gorilla/mux"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(path string) error {
	http.HandleFunc(path, s.handleGetFact)
	switch path {
	case "/cat":
		return http.ListenAndServe(":3060", nil)
	case "/dog":
		return http.ListenAndServe(":3000", nil)
	}
	return nil
}

func (s *ApiServer) handleGetFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetAnimalFact(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
