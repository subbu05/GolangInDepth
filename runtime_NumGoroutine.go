package main

import "runtime"

func main() {
	// Get number of current Go routines
	goRoutinesCount := runtime.NumGoroutine()
	print(goRoutinesCount)

}
