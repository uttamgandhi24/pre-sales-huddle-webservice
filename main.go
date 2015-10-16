package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go AddHandlers(true, ":8080", "server.pem", "server.key")
	go AddHandlers(false, ":8280", "", "")
	wg.Wait()
}
