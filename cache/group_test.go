package cache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGetterFunc(t *testing.T) {
	var f Getter
	f = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := ByteSpace("key")
	if k, _ := f.Get("key"); !reflect.DeepEqual(k, expect) {
		t.Errorf("callback failed!")
	}
}

func TestGroup(t *testing.T) {
	cnts := make(map[string]int, len(db))

	f := GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)

		if v, ok := db[key]; ok {
			cnts[key] += 1
			return []byte(v), nil
		}

		return nil, fmt.Errorf("%s not exist", key)
	})

	g := NewGroup("scores", 1<<10, f)

	for k, v := range db {
		// Load from callback function
		if space, err := g.Get(k); err != nil || space.String() != v {
			t.Fatal("failed to get value of Tom")
		}
		// Cache hit
		if _, err := g.Get(k); err != nil || cnts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	if space, err := g.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", space)
	}
}
