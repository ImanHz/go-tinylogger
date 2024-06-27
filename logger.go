package gotinylogger

import (
	"fmt"
	"os"
	"reflect"
)

type TinyLogger struct {
	hasColor bool
}

// returns a new instance of tiny-logger
func New() *TinyLogger {
	return &TinyLogger{hasColor: true}
}

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
// currently works for map, slice and struct
func (t *TinyLogger) Pretty(title string, v any) {

	titleColor := yellow
	headerColor := green
	headerBold := bold

	if !t.hasColor {
		titleColor = ""
		headerColor = ""
		headerBold = ""
	}

	// getting value, kind and type for reflection
	val := reflect.ValueOf(v)
	kind := val.Kind()
	typ := reflect.TypeOf(v)

	// dereference if v is a pointer
	if kind == reflect.Pointer {
		if typ != nil {
			val = val.Elem()
			kind = val.Kind()
			typ = reflect.TypeOf(val)
		}
	}
	switch kind {

	case reflect.Struct:
		fmt.Printf("%s%s%s%s\n", titleColor, headerBold, title, reset)
		printStruct(&val, &typ, headerColor)

	case reflect.Slice:
		fmt.Printf("%s%s%s%s\n", titleColor, headerBold, title, reset)
		printSlice(&val, &typ, headerColor)

	case reflect.Map:
		fmt.Printf("%s%s%s%s\n", titleColor, headerBold, title, reset)
		printMap(&val, headerColor)

	default:
		t.Info(v)
	}
}

func printSlice(val *reflect.Value, typ *reflect.Type, headerColor string) {

	if val == nil || typ == nil {
		fmt.Println("slice print: value or type is null!")
		return
	}
	fmt.Printf("%s%-4s \t %s (Type=%s) %s\n", headerColor, "Index", "Value", (*typ).String(), reset)

	for i := 0; i < val.Len(); i++ {
		row := fmt.Sprintf("%-4d \t %v ", i, val.Index(i))
		fmt.Println(row)

	}

}

func printMap(val *reflect.Value, headerColor string) {
	fmt.Printf("%s%-10s \t %s%s\n", headerColor, "Key", "Value", reset)

	if val == nil {
		fmt.Println("map print: value is null!")
		return
	}

	rng := val.MapRange()
	for rng.Next() {

		row := fmt.Sprintf("%-10v \t %v ", rng.Key(), rng.Value())
		fmt.Println(row)
	}

}

func printStruct(val *reflect.Value, typ *reflect.Type, headerColor string) {
	if val == nil || typ == nil {
		fmt.Println("struct print: value or type is null!")
		return
	}

	fmt.Printf("%s%-10s \t %s%s\n", headerColor, "Field", "Value", reset)

	fieldNum := val.NumField()
	for i := 0; i < fieldNum; i++ {
		field := val.Field(i)
		name := (*typ).Field(i).Name
		row := fmt.Sprintf("%-10v \t %v ", name, field)
		fmt.Println(row)
	}

}
