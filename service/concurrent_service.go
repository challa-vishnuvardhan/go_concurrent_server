package service

import (
	"go_concurrent_server/db"
	"go_concurrent_server/resource"
	"sync"
)

func ServiceClient(input resource.Request) (resource.Response, error) {
	var result resource.Response
	var wg sync.WaitGroup
	data1 := make(chan string)
	data2 := make(chan string)

	wg.Add(2)

	go func() {
		db.ServiceClient(input, &wg, data1)
	}()

	go func() {
		db.ServiceClientResponse(&wg, data2)
	}()

	result.Message = <-data1 + <-data2
	wg.Wait()
	// Wait for both goroutines to finish
	return result, nil
}
