package util

import (
	"math/rand"
	"net/http"
	"time"
)

func RunEvery(duration time.Duration, foo func()) chan struct{} {
	ticker := time.NewTicker(duration)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				foo()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}

func EnableCors(w *http.ResponseWriter, r *http.Request) bool {
	if r.Method == http.MethodOptions {
		(*w).Header().Set("Access-Control-Allow-Credentials", "true")
		(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		(*w).Header().Set("Access-Control-Allow-Methods", "POST")
		(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
		(*w).Header().Set("Access-Control-Max-Age", "3600")
		(*w).WriteHeader(http.StatusNoContent)
		return true
	}
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	return false
}

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = byte(rand.Intn(26) + 97)
	}
	return string(result)
}
