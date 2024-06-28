package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrinting(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	// * 5 Go Routine running because of wg variable
	for i := 0; i < 5; i++ {
		data := fmt.Sprintf("data %d", i)
		wg.Add(1)
		go doPrinting(&wg, data)
	}

	wg.Wait()
}
