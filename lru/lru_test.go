package lru

import (
	"fmt"
	"testing"
)

type Space struct {
	size uint64
}

func (s Space) Size() uint64 {
	return s.size
}

func TestNewCache(t *testing.T) {
	c := NewCache(0, func(key string, Val Space) {})
	fmt.Println(c)
}
