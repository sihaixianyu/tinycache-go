package cache

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHTTPPool(t *testing.T) {
	NewGroup("scores", 1<<10, GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB ··s·earch ··key]", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}

		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:9999"
	pool := NewHTTPPool(addr)

	log.Println("tinycache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, pool))
}
