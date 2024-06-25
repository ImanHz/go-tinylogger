package gotinylogger

import (
	"fmt"
	"os"
	"reflect"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	bold   = "\033[1m"
)

var hasColor = true

// turns on color mode for macOs/linux terminals
func SetColor(onOff bool) {
	if onOff {
		hasColor = true
	} else {
		hasColor = false
	}
}

// prints error logs
func Error(v ...any) {
	print(red, "ERR", v...)
}

// prints info logs
func Info(v ...any) {
	print(green, "INF", v...)
}

// prints warning logs
func Warning(v ...any) {
	print(yellow, "WRN", v...)
}

// prints fatal log and exits with @code
func Panic(code int, v ...any) {
	print(red, "FTL", v...)
	os.Exit(code)
}

// prints human-readable version of @v. title is the table title
func Pretty(title string, v any) {

	val := reflect.ValueOf(v)
	kind := val.Kind()
	typ := reflect.TypeOf(v)

	switch kind {
	case reflect.Struct:

		fmt.Printf("%s%s%s%s\n", yellow, bold, title, reset)
		fmt.Printf("%s%-10s \t %s%s\n", green, "Field", "Value", reset)

		fieldNum := val.NumField()
		for i := 0; i < fieldNum; i++ {
			field := val.Field(i)
			name := typ.Field(i).Name
			row := fmt.Sprintf("%-10v \t %v ", name, field)
			fmt.Println(row)
		}

	case reflect.Slice:
		fmt.Printf("%s%s%s%s\n", yellow, bold, title, reset)
		fmt.Printf("%s%-4s \t %s (Type=%s) %s\n", green, "Index", "Value", typ.String(), reset)
		for i := 0; i < val.Len(); i++ {
			row := fmt.Sprintf("%-4d \t %v ", i, val.Index(i))
			fmt.Println(row)

		}
	case reflect.Map:

		fmt.Printf("%s%s%s%s\n", yellow, bold, title, reset)
		fmt.Printf("%s%-10s \t %s%s\n", green, "Key", "Value", reset)

		rng := val.MapRange()
		for rng.Next() {

			row := fmt.Sprintf("%-10v \t %v ", rng.Key(), rng.Value())
			fmt.Println(row)
		}

	default:
		Info(v)
	}
}
