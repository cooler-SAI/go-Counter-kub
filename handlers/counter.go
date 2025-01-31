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

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusNoContent) // Возвращаем статус 204 (нет содержимого)
		return
	}

	mutex.Lock()
	counter++
	mutex.Unlock()
	_, err := fmt.Fprintf(w, "Количество запросов: %d", counter)
	if err != nil {
		return
	}
}
