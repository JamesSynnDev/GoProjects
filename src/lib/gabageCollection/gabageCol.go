package gabagecollection

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc : ", mem.Alloc)
	fmt.Println("mem.TotalAlloc : ", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc : ", mem.HeapAlloc)
	fmt.Println("mem.NumGC : ", mem.NumGC)
	fmt.Println("------")
}

func Use_GabaCol() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation Failed!")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation Failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)

}

// GODEBUG=gctrace=1 go runt garbageCol.go
// P.115 Go System Programming
