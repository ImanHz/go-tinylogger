# A very simple logging tool for Go
Based on the standard `log` library with minimum dependencies. 

## Install
Just like other go packages, use: ```go get https://github.com/ImanHz/go-tinylogger```
## How to use

1. Create a new logger instance with New()
2. Set the logger output color is SetColor(bool onOff) (default:true, i.e: use color)
3. Simply use on of the Info(v ...any), Warning(v ...any), Error(v ...any) and Panic(code int, v ...any) functions to log to `stdout`. Panic will terminate the program with exit `code`
4. Use Pretty(v any) to display `struct`, `slice` or `map` data structures in a pretty manner

## Example

```go
package main

import gotinylogger "github.com/ImanHz/go-tinylogger"

func main() {

	// obtain a new instance
    log := gotinylogger.New()

	log.Info("Info")
	log.Warning("Warning")
	log.Error("Error")
	log.Pretty("pretty sample", []string{"first", "second"})
	log.SetColor(false)
	log.Info("no color :(")
	log.Panic(1, "panic!")
}

```

Run: `go run main.go`