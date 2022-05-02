package cache

import (
	"reflect"
	"testing"
	"tinycache-go/cache"
	"tinycache-go/space"
)

func TestGetterFunc_Get(t *testing.T) {
	var f Getter
	f = GetterFunc(func(key string) (cache.Sized, error) {
		return space.BytesSpace(key), nil
	})

	expect := space.BytesSpace("key")
	if k, _ := f.Get("key"); !reflect.DeepEqual(k, expect) {
		t.Errorf("callback failed!")
	}
}
