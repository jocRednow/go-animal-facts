package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetAnimalFact(context.Context) (*AnimalFact, error)
}

type AnimalFactService struct {
	url string
}

func NewAnimalFactService(url string) Service {
	return &AnimalFactService{
		url: url,
	}
}

func (s *AnimalFactService) GetAnimalFact(ctx context.Context) (*AnimalFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fact := &AnimalFact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
