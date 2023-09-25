package db

import (
	"fmt"
	"go_concurrent_server/resource"
	"sync"
)

func ServiceClient(input resource.Request, wg *sync.WaitGroup, data1 chan string) {
	defer wg.Done()
	var clientsMutex sync.Mutex
	clientsMutex.Lock()
	output := "Hi" + " " + input.Message
	clientsMutex.Unlock()
	fmt.Println("1", output)
	data1 <- output
}

func ServiceClientResponse(wg *sync.WaitGroup, data2 chan string) {
	defer wg.Done()
	var clientsMutex sync.Mutex
	clientsMutex.Lock()
	output := " How are you ?"
	clientsMutex.Unlock()
	data2 <- output
}
