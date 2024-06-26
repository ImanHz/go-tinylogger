package gotinylogger

import (
	"fmt"
	"os"
	"reflect"
)

type TinyLogger struct {
	hasColor bool
}

func New() *TinyLogger {
	return &TinyLogger{hasColor: true}
}

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	bold   = "\033[1m"
)

// turns on color mode for macOs/linux terminals
func (t *TinyLogger) SetColor(onOff bool) {
	if onOff {
		t.hasColor = true
	} else {
		t.hasColor = false
	}
}

// prints error logs
func (t *TinyLogger) Error(v ...any) {
	t.print(red, "ERR", v...)
}

// prints info logs
func (t *TinyLogger) Info(v ...any) {
	t.print(green, "INF", v...)
}

// prints warning logs
func (t *TinyLogger) Warning(v ...any) {
	t.print(yellow, "WRN", v...)
}

// prints fatal log and exits with @code
func (t *TinyLogger) Panic(code int, v ...any) {
	t.print(red, "FTL", v...)
	os.Exit(code)
}

// prints human-readable version of @v. title is the table title
func (t *TinyLogger) Pretty(title string, v any) {

	titleColor := yellow
	headerColor := green
	headerBold := bold

	if !t.hasColor {

		titleColor = ""
		headerColor = ""
		headerBold = ""
	}
	val := reflect.ValueOf(v)
	kind := val.Kind()
	typ := reflect.TypeOf(v)

	switch kind {
	case reflect.Struct:

		fmt.Printf("%s%s%s%s\n", titleColor, headerBold, title, reset)
		fmt.Printf("%s%-10s \t %s%s\n", headerColor, "Field", "Value", reset)

		fieldNum := val.NumField()
		for i := 0; i < fieldNum; i++ {
			field := val.Field(i)
			name := typ.Field(i).Name
			row := fmt.Sprintf("%-10v \t %v ", name, field)
			fmt.Println(row)
		}

	case reflect.Slice:
		fmt.Printf("%s%s%s%s\n", titleColor, headerBold, title, reset)
		fmt.Printf("%s%-4s \t %s (Type=%s) %s\n", headerColor, "Index", "Value", typ.String(), reset)
		for i := 0; i < val.Len(); i++ {
			row := fmt.Sprintf("%-4d \t %v ", i, val.Index(i))
			fmt.Println(row)

		}
	case reflect.Map:

		fmt.Printf("%s%s%s%s\n", titleColor, headerBold, title, reset)
		fmt.Printf("%s%-10s \t %s%s\n", headerColor, "Key", "Value", reset)

		rng := val.MapRange()
		for rng.Next() {

			row := fmt.Sprintf("%-10v \t %v ", rng.Key(), rng.Value())
			fmt.Println(row)
		}

	default:
		t.Info(v)
	}
}
