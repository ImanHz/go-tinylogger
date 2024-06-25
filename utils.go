package gotinylogger

import (
	"fmt"
	"time"
)

func print(color string, ident string, v ...any) {

	if !hasColor {
		color = ""
	}

	tim := time.Now().Format(time.DateTime)
	fmt.Printf("%s[%s] %s%s ", color, ident, tim, reset)
	fmt.Println(v...)
}
