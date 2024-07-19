package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	// logger string
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	// BEST PRACTICE !!!!
	defer func(start time.Time) {
		// s.logger
		fmt.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())
	// BEST PRACTICE !!!!

	return s.next.GetCatFact(ctx)
}
