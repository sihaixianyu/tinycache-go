package cache

type Sized interface {
	Size() int
}

type Entry struct {
	Key string
	Val Sized
}

type Cacher interface {
	Set(key string, val Sized)
	Get(key string) (val Sized, ok bool)
	MaxBytes() int
	UsedBytes() int
}
