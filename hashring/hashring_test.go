package hashring

import (
	"strconv"
	"testing"
)

func myHashFn(data []byte) uint32 {
	i, _ := strconv.Atoi(string(data))
	return uint32(i)
}

func TestMap(t *testing.T) {
	ring := New(3, myHashFn)

	ring.Put("2", "4", "6")
	testKeys := map[string]string {
		"2": "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testKeys {
		if ring.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	ring.Put("8")
	testKeys["27"] = "8"

	for k, v := range testKeys {
		if ring.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}
