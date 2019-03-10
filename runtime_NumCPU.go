package main

import "runtime"

func main() {

	// Get number of CPUs available on the system
	cpuCount := runtime.NumCPU()
	print(cpuCount)

}