# A very simple logging tool for Go
Based on standard `log` library with minimum dependencies. 

## How to use

1. Set the logger output color is SetColor(bool onOff) (default:true, i.e: use color)
2. Simply use on of the Info(v ...any), Warning(v ...any), Error(v ...any) and Panic(code int, v ...any) functions to log to `stdout`. Panic will terminate the program with exit `code`
3. Use Pretty(v any) to display `struct`, `slice` or `map` data structures in a pretty manner