package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetAnimalFact(ctx context.Context) (fact *AnimalFact, err error) {
	// BEST PRACTICE !!!!
	defer func(start time.Time) {
		if fact.Facts != nil {
			fmt.Printf("fact=%s err=%v took=%v\n", fact.Facts[0], err, time.Since(start))
		} else {
			fmt.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
		}
	}(time.Now())
	// BEST PRACTICE !!!!

	return s.next.GetAnimalFact(ctx)
}
