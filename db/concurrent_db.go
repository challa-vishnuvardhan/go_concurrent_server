package db

import (
	"go_concurrent_server/resource"
	"sync"
)

func ServiceClient(input resource.Request) (resource.Response, error) {
	var clientsMutex sync.Mutex
	clientsMutex.Lock()
	output := "HI" + " " + input.Message
	clientsMutex.Unlock()

	var result resource.Response

	result.Message = output
	return result, nil

}
