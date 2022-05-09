package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	counter := 0
	gs := 10

	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	wg.Add(gs)

	for i := 1; i <= gs; i++ {
		go func() {
			defer wg.Done()
			v := counter
			runtime.Gosched()
			v++
			counter = v
		}()
	}

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())

	fmt.Println("count", counter)
}
