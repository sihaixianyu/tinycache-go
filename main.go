package main

import (
	"fmt"
)

type Space struct {
	size uint64
}

func main() {
	var s Space

	fmt.Printf("%v", s)
}
