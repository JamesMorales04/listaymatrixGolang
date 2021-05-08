package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	var matrix [1000][1000]int

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			matrix[i][j] = rand.Intn(1000)
			go memoryRead()
		}
		
	}
	memoryRead()

}

func memoryRead() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\nAlloc = %v\nTotalAlloc = %v\nSys = %v\nNumGC = %v\n\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
