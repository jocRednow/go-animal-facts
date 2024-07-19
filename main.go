package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	animals := map[string]string{
		"/dog": "https://dog-api.kinduff.com/api/facts?number=1",
		"/cat": "https://catfact.ninja/fact",
	}

	for i, v := range animals {
		wg.Add(1)
		go func() {
			defer wg.Done()
			srv := NewAnimalFactService(v)
			srv = NewLoggingService(srv)
			apiServer := NewApiServer(srv)
			log.Fatal(apiServer.Start(i))
		}()
	}
	wg.Wait()
}
