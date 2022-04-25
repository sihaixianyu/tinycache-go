package util

import (
	"fmt"
	"strings"
)

type Debuger interface {
	Format(level int) string
}

func InsertTab(s *strings.Builder, level int) {
	for i := 0; i < level; i++ {
		s.WriteString("\t")
	}
}

func Debug(d Debuger) {
	fmt.Println(d.Format(0))
}
