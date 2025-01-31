package handlers

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func CounterHandler(w http.ResponseWriter, _ *http.Request) {

	mutex.Lock()
	counter++
	mutex.Unlock()
	_, err := fmt.Fprintf(w, "Counter: %d", counter)
	if err != nil {
		return
	}
}
