package gotinylogger

import (
	"fmt"
	"time"
)

func (t *TinyLogger) print(color string, ident string, v ...any) {

	if !t.hasColor {
		color = ""
	}

	tim := time.Now().Format(time.DateTime)
	fmt.Printf("%s[%s] %s%s ", color, ident, tim, reset)
	fmt.Println(v...)
}
